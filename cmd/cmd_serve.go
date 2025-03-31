package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/cocoth/linknet-api/src/controllers"
	"github.com/cocoth/linknet-api/src/database"
	"github.com/cocoth/linknet-api/src/http/middleware"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/cocoth/linknet-api/src/routes"
	"github.com/cocoth/linknet-api/src/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

var (
	appDb        string
	appHost      string
	appPort      string
	appUpload    string
	dbName       string
	dbHost       string
	dbPort       string
	dbUser       string
	dbPass       string
	dbKeyEncrypt string
	jwtSecret    string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		ServerConfig()
		InitializeAndRunServer()
	},
}

func init() {
	serveCmd.Flags().StringVar(&appDb, "dbms", "", "database management system (e.g. mysql, postgres)")
	serveCmd.Flags().StringVar(&appHost, "host", "", "Application host")
	serveCmd.Flags().StringVar(&appPort, "port", "", "Application port")
	serveCmd.Flags().StringVar(&appUpload, "upload", "", "Upload directory")
	serveCmd.Flags().StringVar(&dbName, "db-name", "", "Database name")
	serveCmd.Flags().StringVar(&dbHost, "db-host", "", "Database host")
	serveCmd.Flags().StringVar(&dbPort, "db-port", "", "Database port")
	serveCmd.Flags().StringVar(&dbUser, "db-user", "", "Database user")
	serveCmd.Flags().StringVar(&dbPass, "db-pass", "", "Database password")
	serveCmd.Flags().StringVar(&dbKeyEncrypt, "db-key-encrypt", "", "Database key encrypt for user password, must be 16, 24, or 32 bytes")
	serveCmd.Flags().StringVar(&jwtSecret, "jwt-secret", "", "JWT secret")
	RootCmd.AddCommand(serveCmd)
}

func ServerConfig() {
	if appDb == "" {
		appDb = os.Getenv("APP_DB")
		if appDb == "" {
			appDb = PromptInput("Enter database management system (e.g. mysql, postgres)")
			UpdateEnv("APP_DB", appDb)
		}
	} else {
		UpdateEnv("APP_DB", appDb)
	}
	if appHost == "" {
		appHost = os.Getenv("APP_HOST")
		if appHost == "" {
			appHost = PromptInput("Enter application host")
			UpdateEnv("APP_HOST", appHost)
		}
	} else {
		UpdateEnv("APP_HOST", appHost)
	}
	if appPort == "" {
		appPort = os.Getenv("APP_PORT")
		if appPort == "" {
			appPort = PromptInput("Enter application port")
			UpdateEnv("APP_PORT", appPort)
		}
	} else {
		UpdateEnv("APP_PORT", appPort)
	}
	if appUpload == "" {
		appUpload = os.Getenv("APP_UPLOAD_DIR")
		if appUpload == "" {
			appUpload = PromptInput("Enter upload directory")
			UpdateEnv("APP_UPLOAD", appUpload)
		}
	} else {
		UpdateEnv("APP_UPLOAD_DIR", appUpload)
	}
	if dbName == "" {
		dbName = os.Getenv("DB_NAME")
		if dbName == "" {
			dbName = PromptInput("Enter database name")
			UpdateEnv("DB_NAME", dbName)
		}
	} else {
		UpdateEnv("DB_NAME", dbName)
	}
	if dbHost == "" {
		dbHost = os.Getenv("DB_HOST")
		if dbHost == "" {
			dbHost = PromptInput("Enter database host")
			UpdateEnv("DB_HOST", dbHost)
		}
	} else {
		UpdateEnv("DB_HOST", dbHost)
	}
	if dbPort == "" {
		dbPort = os.Getenv("DB_PORT")
		if dbPort == "" {
			dbPort = PromptInput("Enter database port")
			UpdateEnv("DB_PORT", dbPort)
		}
	} else {
		UpdateEnv("DB_PORT", dbPort)
	}
	if dbUser == "" {
		dbUser = os.Getenv("DB_USER")
		if dbUser == "" {
			dbUser = PromptInput("Enter database user")
			UpdateEnv("DB_USER", dbUser)
		}
	} else {
		UpdateEnv("DB_USER", dbUser)
	}
	if dbPass == "" {
		dbPass = os.Getenv("DB_PASS")
		if dbPass == "" {
			dbPass = PromptInputCredentials("Enter database password")
			UpdateEnv("DB_PASS", dbPass)
		}
	} else {
		UpdateEnv("DB_PASS", dbPass)
	}
	if dbKeyEncrypt == "" {
		dbKeyEncrypt = os.Getenv("DB_KEY_ENCRYPT")
		if dbKeyEncrypt == "" {
			for {
				dbKeyEncrypt = PromptInputCredentials("Enter database key encrypt for user password, min 16 bytes, 24 bytes, or 32 bytes")
				if len(dbKeyEncrypt) != 16 && len(dbKeyEncrypt) != 24 && len(dbKeyEncrypt) != 32 {
					fmt.Println("Encryption key must be at least min 16 bytes, 24 bytes, or 32 bytes. Please try again.")
					continue
				}
				break
			}
			UpdateEnv("DB_KEY_ENCRYPT", dbKeyEncrypt)
		}
	} else {
		UpdateEnv("DB_KEY_ENCRYPT", dbKeyEncrypt)
	}
	if jwtSecret == "" {
		jwtSecret = os.Getenv("JWT_SECRET_KEY_USER")
		if jwtSecret == "" {
			jwtSecret = PromptInput("Enter JWT secret")
			UpdateEnv("JWT_SECRET_KEY_USER", jwtSecret)
		}
	} else {
		UpdateEnv("JWT_SECRET_KEY_USER", jwtSecret)
	}
}

func InitializeAndRunServer() {
	db := database.DB
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Disposition", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition", "Set-Cookie"},
		AllowWebSockets:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	validate := validator.New()

	v1 := r.Group("/api/v1")

	userRepo := repo.NewUserRepoImpl(db)
	fileUploadRepo := repo.NewFileUploadRepoImpl(db)
	surveyRepo := repo.NewSurveyRepoImpl(db)
	filePermRepo := repo.NewFilePermRepoImpl(db)
	notifRepo := repo.NewNotifyRepoImpl(db)

	userService := services.NewUserServiceImpl(userRepo, validate)
	authService := services.NewAuthService(userRepo)
	fileUploadService := services.NewFileUploadServiceImpl(fileUploadRepo)
	filePermService := services.NewFilePermServceImpl(filePermRepo)
	surveyService := services.NewSurveyServiceImpl(surveyRepo)
	notifService := services.NewNotifyServiceImpl(notifRepo)

	authMiddleware := middleware.NewUserAuthorization(userService)

	userCtrl := controllers.NewUserController(userService)
	authCtrl := controllers.NewAuthController(authService)
	fileCtrl := controllers.NewFileController(fileUploadService, filePermService, userService)
	surveyCtrl := controllers.NewSurveyController(surveyService, userService)
	notifCtrl := controllers.NewNotifyController(notifService)

	routes.AuthRoute(authMiddleware, authCtrl, v1)
	routes.UserRoute(userCtrl, v1)
	routes.FileRoute(fileCtrl, v1)
	routes.SurveyRoute(surveyCtrl, v1)
	routes.NotificationRoute(notifCtrl, v1)

	StartServer(r, appHost, appPort)
}
