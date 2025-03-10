package models

import (
	"time"

	"github.com/cocoth/linknet-api/src/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id         string `gorm:"type:uuid;primaryKey" json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Email      string `gorm:"unique;not null" json:"email"`
	Password   string `json:"password"`
	RoleID     uint   `json:"role_id"`
	Role       Role   `gorm:"foreignKey:RoleID" json:"role"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Deleted_at time.Time `gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Id = uuid.New().String()
	hash, err := utils.GenerateHashPassword([]byte(user.Password))
	user.Password = hash
	if err != nil {
		panic(err)
	}
	return nil
}

type Role struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type Permission struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}
