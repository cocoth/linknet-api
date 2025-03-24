package models

import (
	"time"

	"github.com/cocoth/linknet-api/src/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         string       `gorm:"type:varchar(36);primaryKey" json:"id"`
	Name       string       `gorm:"not null" json:"name"`
	Email      string       `gorm:"unique;not null" json:"email"`
	Phone      string       `gorm:"unique;not null" json:"phone"`
	Password   string       `json:"password"`
	CallSign   string       `json:"call_sign"`
	Contractor *string      `json:"contractor"`
	Status     *string      `gorm:"type:enum('active','inactive');default:'active'" json:"status"`
	RoleID     *uint        `json:"role_id"`
	Role       *Role        `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	FileUpload []FileUpload `gorm:"foreignKey:AuthorID" json:"file_upload"`
	Notifies   []Notify     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"notifies"`
	Surveys    []Survey     `gorm:"foreignKey:SurveyorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"surveys"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
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
