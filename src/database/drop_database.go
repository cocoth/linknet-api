package database

import (
	"os"

	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/utils"
)

func DropTables() {
	err := DB.Migrator().DropTable(
		&models.Role{},
		&models.User{},
		&models.FileUpload{},
		&models.FileAccessRequest{},
		&models.Notify{},
		&models.Survey{},
		&models.SurveyorLink{},
		&models.SurveyReport{},
		&models.ISmart{},
	)
	if err != nil {
		utils.Error("Failed to drop tables", "DropTables")
		os.Exit(1)
	}
}
