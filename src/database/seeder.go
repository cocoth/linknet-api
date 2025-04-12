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

	// Skip the header row
	if len(records) > 0 {
		records = records[1:] // Remove the header row
	}

	// Define batch size
	batchSize := 500
	totalRecords := len(records)

	// Channel to track completion of all goroutines
	done := make(chan bool, (totalRecords+batchSize-1)/batchSize)

	// Process records in batches
	for i := 0; i < totalRecords; i += batchSize {
		end := min(i+batchSize, totalRecords)

		batch := records[i:end]

		// Use a goroutine to process each batch
		go func(batch [][]string) {
			for _, record := range batch {
				if len(record) < 3 {
					log.Printf("Skipping invalid record: %v", record)
					continue
				}

				// Map CSV fields to ISmart model
				ismart := models.ISmart{
					FiberNode:  record[0],
					Address:    record[1],
					Coordinate: record[2],
					Street:     "", // Default value for Street
				}

				// Check if the record already exists
				var existing models.ISmart
				DB.First(&existing, "fiber_node = ?", ismart.FiberNode)
				if existing.ID == "" {
					// Create a new record if it doesn't exist
					DB.Create(&ismart)
					log.Printf("Added to database: %+v", ismart) // Output to terminal
				}
			}
			done <- true // Signal completion of this batch
		}(batch)
	}

	// Wait for all goroutines to finish
	for range (totalRecords + batchSize - 1) / batchSize {
		<-done
	}

	log.Println("done add all i smart")
}
