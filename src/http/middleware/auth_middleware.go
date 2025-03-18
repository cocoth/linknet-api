package middleware

import (
	"time"

	ctrlUtils "github.com/cocoth/linknet-api/src/controllers/utils"
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

	user, err := u.authService.GetById(userId)
	if err != nil {
		ctrlUtils.RespondWithError(c, 401, "Unauthorized")
		c.Abort()
		return
	}

	// c.Set("X-CSRF-Token", csrfToken)
	c.Set("user", user)
	c.Next()
}
