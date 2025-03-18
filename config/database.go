package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/cocoth/linknet-api/config/models"
	"github.com/cocoth/linknet-api/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func createDatabaseIfNotExists(driver, dsn, dbName string) error {
	var db *sql.DB
	var err error

	if driver == "mysql" {
		db, err = sql.Open("mysql", dsn)
	} else if driver == "postgres" {
		db, err = sql.Open("postgres", dsn)
	}

	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	return err
}

func ConnectToDB() {
	var err error
	var dsn string

	var (
		appDB   = os.Getenv("APP_DB")
		db_host = os.Getenv("DB_HOST")
		db_port = os.Getenv("DB_PORT")
		db_user = os.Getenv("DB_USER")
		db_name = os.Getenv("DB_NAME")
		db_pass = os.Getenv("DB_PASS")
	)

	if appDB == "mysql" {
		dsn = db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/"
		err = createDatabaseIfNotExists("mysql", dsn, db_name)
		if err != nil {
			utils.Error("Failed to create database", "ConnectToDB")
			utils.ErrPanic(err)
		}
		dsn += db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if appDB == "postgres" {
		dsn = "host=" + db_host + " port=" + db_port + " user=" + db_user + " password=" + db_pass + " sslmode=disable"
		err = createDatabaseIfNotExists("postgres", dsn, db_name)
		if err != nil {
			utils.Error("Failed to create database", "ConnectToDB")
			utils.ErrPanic(err)
		}
		dsn += " dbname=" + db_name
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		utils.Error("Failed to connect to database", "ConnectToDB")
		utils.ErrPanic(err)
	}

	DB.AutoMigrate(&models.User{}, &models.Role{})
}

func RoleSeeder() {
	roles := []models.Role{
		{Name: "user"},
		{Name: "admin"},
	}
	for _, role := range roles {
		var existing models.Role
		DB.First(&existing, "name = ?", role.Name)
		if existing.ID == 0 {
			DB.Create(&role)
		}
	}
}

func DropTables() {
	err := DB.Migrator().DropTable(&models.User{}, &models.Role{})
	if err != nil {
		utils.Error("Failed to drop tables", "DropTables")
		utils.ErrPanic(err)
	}
}
