package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/cocoth/linknet-api/src/database"
	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/repo"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var initdbCmd = &cobra.Command{
	Use:     "initdb",
	Short:   "Initialize the database",
	Long:    "Run this command once to initialize the database, setup default admin user and setup default admin password.\nThis command is a must when running the application for the first time.",
	Example: `linknet-api initdb`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("====================================")
		fmt.Println("Initializing the database...")
		fmt.Println("====================================")
		InitializeDatabase(database.DB)
		fmt.Println("====================================")
		fmt.Println("Database initialization complete!")
		fmt.Println("====================================")
	},
}

func init() {
	RootCmd.AddCommand(initdbCmd)
}

func InitializeDatabase(db *gorm.DB) {
	database.DropTables()
	database.ConnectToDB()
	database.RoleSeeder()
	database.ISmartSeeder("config/i-smart.csv")

	email := PromptInput("Enter new admin email")

	phone := PromptInput("Enter new admin phone")

	var password, retypePassword string
	for {
		password = PromptInputCredentials("Enter new admin password")

		if len(password) < 8 {
			fmt.Println("Password must be at least 8 characters long. Please try again.")
			continue
		}

		retypePassword = PromptInputCredentials("Retype new admin password")

		if password == retypePassword {
			break
		} else {
			fmt.Println("Passwords do not match. Please try again.")
		}
	}

	adminUser := models.User{
		Name:     "admin",
		Email:    email,
		Phone:    phone,
		Password: retypePassword,
	}

	repo := repo.NewUserRepoImpl(db)
	role := "admin"

	roleDb, err := repo.GetRoleByRoleName(role)
	if err != nil {
		fmt.Println("Failed to get role")
		fmt.Println(err)
		os.Exit(1)
	}
	adminUser.Role = &roleDb
	adminUser.RoleID = &roleDb.ID

	adminDbResult, err := repo.CreateUser(adminUser)
	if err != nil {
		fmt.Println("Failed to create admin user")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Admin user created successfully!")
	fmt.Println("Please login with email: ", adminDbResult.Email, " and password: ", retypePassword)

	time.Sleep(2 * time.Second)
	InitializeAndRunServer()
}
