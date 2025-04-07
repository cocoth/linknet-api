package cmd

import (
	"os"

	"github.com/cocoth/linknet-api/src/utils"
	"github.com/gin-gonic/gin"
)

func StartServer(e *gin.Engine, appHost, appPort string) {
	if appHost == "" {
		appHost = os.Getenv("APP_HOST")
		if appHost == "" {
			appHost = "0.0.0.0"
		}
	}

	if appPort == "" {
		appPort = os.Getenv("APP_PORT")
		if appPort == "" {
			appPort = "3000"
		}

	}

	env := os.Getenv("APP_ENV")

	if env == "prod" {
		os.Setenv("GIN_MODE", gin.ReleaseMode)
		gin.SetMode(gin.ReleaseMode)
	}

	serverAddr := appHost + ":" + appPort
	if err := e.Run(serverAddr); err != nil {
		utils.Error(err.Error(), "RunServer")
		os.Exit(1)
	}
}
