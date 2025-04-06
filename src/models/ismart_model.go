package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ISmart struct {
	ID         string `gorm:"type:varchar(36);primaryKey" json:"id"`
	FiberNode  string `json:"fiber_node"`
	Address    string `json:"address"`
	Coordinate string `json:"coordinate"`
	Street     string `json:"street"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}

func (i *ISmart) BeforeCreate(tx *gorm.DB) error {
	i.ID = uuid.New().String()
	return nil
}
