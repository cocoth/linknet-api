package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileUpload struct {
	ID        string  `gorm:"type:varchar(36);primaryKey" json:"id"`
	FileName  string  `gorm:"not null" json:"file_name"`
	FileUri   string  `json:"file_uri"`
	FileHash  string  `json:"file_hash"`
	AuthorID  *string `json:"author_id"`
	Author    *User   `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (user *FileUpload) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	return nil
}
