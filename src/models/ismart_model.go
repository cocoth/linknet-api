package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ISmart struct {
	ID           string `gorm:"type:varchar(36);primaryKey" json:"id"`
	FiberNode    string `json:"fiber_node"`
	Address      string `json:"address"`
	CustomerName string `json:"customer_name"`
	Coordinate   string `json:"coordinate"`
	Street       string `json:"street"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `gorm:"index"`
}

func (s *ISmart) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New().String()
	return nil
}
