package database

import (
	"log"

	"github.com/cocoth/linknet-api/src/models"
	"github.com/cocoth/linknet-api/src/utils"
)

func RoleSeeder() {
	roles := []models.Role{
		{Name: "admin"},
		{Name: "user"},
	}
	for _, role := range roles {
		var existing models.Role
		DB.First(&existing, "name = ?", role.Name)
		if existing.ID == 0 {
			DB.Create(&role)
		}
	}
}

func ISmartSeeder(filePath string) {
	// Read data from the CSV file
	records, err := utils.ReadCSVFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Skip the header row and iterate over the records
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		if len(record) < 3 {
			log.Printf("Skipping invalid record at line %d: %v", i+1, record)
			continue
		}

		// Handle empty Street field
		street := record[3]
		if street == "" {
			street = "" // Set to an empty string if the field is empty
		}

		// Map CSV fields to ISmart model
		ismart := models.ISmart{
			FiberNode:  record[0],
			Address:    record[1],
			Coordinate: record[2],
			Street:     street,
		}

		// Check if the record already exists
		var existing models.ISmart
		DB.First(&existing, "fiber_node = ?", ismart.FiberNode)
		if existing.ID == "" {
			// Create a new record if it doesn't exist
			DB.Create(&ismart)
		}
	}
}
