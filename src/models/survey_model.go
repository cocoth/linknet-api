package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SurveyorLink struct {
	ID         string `gorm:"type:varchar(36);primaryKey" json:"id"`
	SurveyID   string `json:"survey_id"`
	SurveyorID string `json:"surveyor_id"`
	Surveyor   User   `gorm:"foreignKey:SurveyorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"surveyor"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}

func (s *SurveyorLink) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New().String()
	return nil
}

type Survey struct {
	ID           string         `gorm:"type:varchar(36);primaryKey" json:"id"`
	Title        string         `gorm:"not null" json:"title"`
	FormNumber   string         `json:"form_number"`
	QuestorName  string         `json:"questor_name"`
	FAT          string         `json:"fat"`
	CustomerName string         `json:"customer_name"`
	Address      string         `json:"address"`
	NodeFDT      string         `json:"node_fdt"`
	SurveyDate   time.Time      `json:"survey_date"`
	Surveyors    []SurveyorLink `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"surveyors"`
	ImageID      *string        `json:"image_id"`
	Image        *FileUpload    `gorm:"foreignKey:ImageID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"image"`
	SurveyReport *SurveyReport  `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"survey_report"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `gorm:"index"`
}

func (s *Survey) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New().String()
	return nil
}

type SurveyReport struct {
	ID        string      `gorm:"type:varchar(36);primaryKey" json:"id"`
	SurveyID  string      `json:"survey_id"`
	Survey    Survey      `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Remark    string      `gorm:"type:text" json:"remark"`
	Status    string      `gorm:"type:varchar(20);default:'incomplete'" json:"status"`
	ImageID   *string     `gorm:"type:varchar(36)" json:"image_id"`
	Image     *FileUpload `gorm:"foreignKey:ImageID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (s *SurveyReport) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New().String()
	return nil
}

type ISmart struct {
	ID           string `gorm:"type:varchar(36);primaryKey" json:"id"`
	FiberNode    string `json:"fiber_node"`
	Address      string `json:"address"`
	CustomerName string `json:"customer_name"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `gorm:"index"`
}

func (s *ISmart) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New().String()
	return nil
}

type Notify struct {
	ID            string      `gorm:"type:varchar(36);primaryKey" json:"id"`
	UserID        string      `json:"user_id"`
	User          User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	FileID        *string     `json:"file_id"`
	File          *FileUpload `gorm:"foreignKey:FileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"file"`
	NotifyType    string      `json:"notify_type"`
	NotifyStatus  string      `gorm:"type:varchar(20);default:'pending'" json:"notify_status"`
	NotifyMessage string      `json:"notify_message"`
	IsRead        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `gorm:"index"`
}

func (n *Notify) BeforeCreate(tx *gorm.DB) error {
	n.ID = uuid.New().String()
	return nil
}
