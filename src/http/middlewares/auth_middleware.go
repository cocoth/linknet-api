package middlewares

import (
	"os"
	"strings"
	"time"

	"github.com/cocoth/linknet-api/src/controllers/helper"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/cocoth/linknet-api/src/utils"
	"github.com/gin-gonic/gin"
)

type UserAuthorization struct {
	authService services.UserService
}

func NewUserAuthorization(service services.UserService) *UserAuthorization {
	return &UserAuthorization{
		authService: service,
	}
}

func (u *UserAuthorization) Authorize(c *gin.Context) {
	var sessionToken string

	authHeader := c.GetHeader("Authorization")
	utils.Debug("Auth Header: " + authHeader)

	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		sessionToken = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		var err error
		sessionToken, err = c.Cookie("session_token")
		if err != nil || sessionToken == "" {
			helper.RespondWithError(c, 401, "Unauthorized")
			c.Abort()
			return
		}
	}

	exp, userId, err := utils.ValidateJWTToken(sessionToken)

	if err != nil {
		helper.RespondWithError(c, 401, "Invalid Token")
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > exp {
		helper.RespondWithError(c, 401, "Token expired")
		c.Abort()
		return
	}

	user, err := u.authService.GetUserById(userId)
	if err != nil {
		helper.RespondWithError(c, 401, "Unauthorized")
		c.Abort()
		return
	}

	domain := os.Getenv("APP_DOMAIN")

	c.Set("current_user", user)
	c.SetCookie("session_token", sessionToken, int((24 * time.Hour).Seconds()), "/", domain, false, true)

	c.Next()
}
