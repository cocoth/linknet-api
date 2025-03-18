package models

import "time"

type FileUpload struct {
	ID         string `gorm:"type:varchar(36);primaryKey" json:"id"`
	FileName   string `gorm:"not null" json:"file_name"`
	FileUri    string `json:"file_uri"`
	FileHash   string `json:"file_hash"`
	AuthorID   string `json:"author_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Deleted_at *time.Time `gorm:"index"`
}
