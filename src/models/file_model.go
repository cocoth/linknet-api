package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileUpload struct {
	ID        string  `gorm:"type:varchar(36);primaryKey" json:"id"`
	FileName  string  `gorm:"not null" json:"file_name"`
	FileType  string  `gorm:"not null" json:"file_type"`
	FileUri   string  `json:"file_uri"`
	FileHash  string  `json:"file_hash"`
	AuthorID  *string `json:"author_id"`
	Author    *User   `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (f *FileUpload) BeforeCreate(tx *gorm.DB) error {
	f.ID = uuid.New().String()
	return nil
}

type FileAccessRequest struct {
	ID         string      `gorm:"type:varchar(36);primaryKey" json:"id"`
	FileID     string      `json:"file_id"`
	FileUpload *FileUpload `gorm:"foreignKey:FileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"file_upload"`
	UserID     string      `json:"requestor_id"`
	User       *User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"requestor"`
	Approved   bool        `json:"approved"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}

func (f *FileAccessRequest) BeforeCreate(tx *gorm.DB) error {
	f.ID = uuid.New().String()
	return nil
}

type FileAccess struct {
	ID         string      `gorm:"type:varchar(36);primaryKey" json:"id"`
	FileID     string      `json:"file_id"`
	FileUpload *FileUpload `gorm:"foreignKey:FileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"file_upload"`
	UserID     string      `json:"requestor_id"`
	User       *User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"requestor"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}

func (f *FileAccess) BeforeCreate(tx *gorm.DB) error {
	f.ID = uuid.New().String()
	return nil
}
