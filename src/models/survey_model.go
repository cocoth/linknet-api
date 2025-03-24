package models

import "time"

type Survey struct {
	ID           string        `gorm:"type:varchar(36);primaryKey" json:"id"`
	Title        string        `gorm:"not null" json:"title"`
	FormNumber   string        `json:"form_number"`
	QuestorName  string        `json:"questor_name"`
	FAT          string        `json:"fat"`
	CustomerName string        `json:"customer_name"`
	Address      string        `json:"address"`
	NodeFDT      string        `json:"node_fdt"`
	SurveyDate   time.Time     `json:"survey_date"`
	SurveyorID   string        `json:"surveyor_id"`
	Surveyor     User          `gorm:"foreignKey:SurveyorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"surveyor"`
	ImageID      *string       `json:"image_id"`
	Image        *FileUpload   `gorm:"foreignKey:ImageID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"image"`
	SurveyReport *SurveyReport `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"survey_report"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `gorm:"index"`
}

type SurveyReport struct {
	ID        string      `gorm:"type:varchar(36);primaryKey" json:"id"`
	SurveyID  string      `json:"survey_id"`
	Survey    Survey      `gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"survey"`
	Remark    string      `json:"remark"`
	Status    string      `gorm:"type:enum('standard','reject','incomplete');default:'incomplete'" json:"status"`
	ImageID   *string     `json:"image_id"`
	Image     *FileUpload `gorm:"foreignKey:ImageID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

type Notify struct {
	ID            string      `gorm:"type:varchar(36);primaryKey" json:"id"`
	UserID        string      `gorm:"type:varchar(36);not null" json:"user_id"`
	User          User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	FileID        string      `json:"file_id"`
	File          *FileUpload `gorm:"foreignKey:FileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"file"`
	NotifyType    string      `json:"notify_type"`
	NotifyStatus  string      `gorm:"type:enum('pending','approved','rejected');default:'pending'" json:"notify_status"`
	NotifyMessage string      `json:"notify_message"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `gorm:"index"`
}
