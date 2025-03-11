package models

import (
	"errors"
	"time"

	"github.com/cocoth/linknet-api/src/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id         string `gorm:"type:uuid;primaryKey" json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Email      string `gorm:"unique;not null" json:"email"`
	Phone      string `gorm:"unique;not null" json:"phone"`
	Password   string `json:"password"`
	RoleID     *uint  `json:"role_id"`
	Role       *Role  `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Deleted_at time.Time `gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Id = uuid.New().String()
	hash, err := utils.GenerateHashPassword([]byte(user.Password))
	if err != nil {
		panic(err)
	}
	user.Password = hash
	return nil
}

type Role struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Users []User `gorm:"foreignKey:RoleID" json:"users"`
}

func (role *Role) BeforeSave(tx *gorm.DB) error {
	listRoles := []string{"user", "admin"}
	for _, r := range listRoles {
		if role.Name == r {
			return nil
		}
	}
	return errors.New("invalid role name")
}
