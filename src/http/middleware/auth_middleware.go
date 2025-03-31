package middleware

import (
	"os"
	"time"

	ctrlUtils "github.com/cocoth/linknet-api/src/controllers/helper"
	"github.com/cocoth/linknet-api/src/services"
	generalUtils "github.com/cocoth/linknet-api/src/utils"
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
	sessionToken, errSess := c.Cookie("session_token")

	if errSess != nil {
		ctrlUtils.RespondWithError(c, 401, "Unauthorized")
		c.Abort()
		return
	}
	// csrfToken := c.GetHeader("X-CSRF-Token")
	// generateCSRF := generalUtils.GenerateCSRFToken(32)

	// generalUtils.Debug(csrfToken)
	// generalUtils.Debug(generateCSRF)
	// if csrfToken == "" || csrfToken != generateCSRF {
	// 	ctrlUtils.RespondWithError(c, 401, "PPPPPPP")
	// 	c.Abort()
	// 	return
	// }

	exp, userId, err := generalUtils.ValidateJWTToken(sessionToken)

	if err != nil {
		ctrlUtils.RespondWithError(c, 401, "Invalid Token")
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > exp {
		ctrlUtils.RespondWithError(c, 401, "Token expired")
		c.Abort()
		return
	}

	user, err := u.authService.GetUserById(userId)
	if err != nil {
		ctrlUtils.RespondWithError(c, 401, "Unauthorized")
		c.Abort()
		return
	}

	// c.Set("X-CSRF-Token", csrfToken)
	c.Set("user", user)
	domain := os.Getenv("APP_DOMAIN")
	c.SetCookie("session_token", sessionToken, int((24 * time.Hour).Seconds()), "/", domain, false, true)

	c.Next()
}
