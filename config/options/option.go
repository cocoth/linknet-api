package options

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/cocoth/linknet-api/config"
	"github.com/cocoth/linknet-api/config/models"
	"github.com/cocoth/linknet-api/src/repo"
	"gorm.io/gorm"
)

func Opt(db *gorm.DB) {
	dropDB := flag.Bool("drop-db", false, "Drop database and create new one")
	init := flag.Bool("init", false, "initialize database")

	flag.Parse()

	if db == nil {
		fmt.Println("Database connection is nil")
		os.Exit(1)
	}

	if *dropDB {
		fmt.Print("Flushing Database...\n\n")
		time.Sleep(1 * time.Second)
		config.DropTables()
		config.ConnectToDB()
	}
	if *init {
		config.DropTables()
		config.ConnectToDB()
		config.RoleSeeder()

		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter new admin email: ")
		email, _ := r.ReadString('\n')
		email = email[:len(email)-1]

		fmt.Print("Enter new admin phone: ")
		phone, _ := r.ReadString('\n')
		phone = phone[:len(phone)-1]

		var password, retypePassword string
		for {
			fmt.Print("Enter new admin password: ")
			password, _ = r.ReadString('\n')
			password = password[:len(password)-1]

			if len(password) < 8 {
				fmt.Println("Password must be at least 8 characters long. Please try again.")
				continue
			}

			fmt.Print("Retype new admin password: ")
			retypePassword, _ = r.ReadString('\n')
			retypePassword = retypePassword[:len(retypePassword)-1]

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

		// err := db.Where("name = ?", role).First(&role).Error
		roleDb, err := repo.GetRoleByRoleName(role)
		if err != nil {
			fmt.Println("Failed to get role")
			fmt.Println(err)
			os.Exit(1)
		}
		adminUser.Role = &roleDb
		adminUser.RoleID = &roleDb.ID

		adminDbResult, err := repo.Create(adminUser)
		if err != nil {
			fmt.Println("Failed to create admin user")
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Admin user created successfully")
		fmt.Println("please login with email: ", adminDbResult.Email, " and password: ", retypePassword)

	}
}
