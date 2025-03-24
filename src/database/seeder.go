package database

import "github.com/cocoth/linknet-api/src/models"

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
