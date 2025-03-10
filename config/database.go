package config

import (
	"os"

	"github.com/cocoth/linknet-api/config/models"
	"github.com/cocoth/linknet-api/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
		dsn = db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if appDB == "postgres" {
		dsn = "host=" + db_host + " port=" + db_port + " user=" + db_user + " dbname=" + db_name + " password=" + db_pass + " sslmode=disable"
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		utils.Error("Failed to connect to database", "ConnectToDB")
		utils.ErrPanic(err)
	}

	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.RolePermission{})
}

func DropTables() {
	err := DB.Migrator().DropTable(&models.User{}, &models.Role{}, &models.Permission{}, &models.RolePermission{})
	if err != nil {
		utils.Error("Failed to drop tables", "DropTables")
		utils.ErrPanic(err)
	}
}
