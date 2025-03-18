package models

import "time"

type Survey struct {
	ID           string     `gorm:"type:varchar(36);primaryKey" json:"id"`
	Title        string     `gorm:"not null" json:"title"`
	FormNumber   string     `json:"form_number"`
	QuestorName  string     `json:"questor_name"`
	FAT          string     `json:"fat"`
	CustomerName string     `json:"customer_name"`
	Address      string     `json:"address"`
	NodeFDT      string     `json:"node_fdt"`
	SurveyDate   time.Time  `json:"survey_date"`
	SurveyorID   string     `json:"surveyor_id"`
	ImageID      *string    `json:"image_id"`
	Image        FileUpload `gorm:"foreignKey:ImageID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"image"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Deleted_at   *time.Time `gorm:"index"`
}

type SurveyReport struct {
	ID         string     `gorm:"type:varchar(36);primaryKey" json:"id"`
	SurveyID   string     `json:"survey_id"`
	Remark     string     `json:"remark"`
	Status     string     `gorm:"type:enum('standard','reject','incomplete');default:'incomplete'" json:"status"`
	ImageID    *string    `json:"image_id"`
	Image      FileUpload `gorm:"foreignKey:ImageID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"image"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Deleted_at *time.Time `gorm:"index"`
}

type Notify struct {
	ID            string `gorm:"type:varchar(36);primaryKey" json:"id"`
	NotifyType    string `json:"notify_type"`
	NotifyID      string `json:"notify_id"`
	NotifyFrom    string `json:"notify_from"`
	NotifyTo      string `json:"notify_to"`
	NotifyMessage string `json:"notify_message"`
	NotifyStatus  string `json:"notify_status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Deleted_at    *time.Time `gorm:"index"`
}
