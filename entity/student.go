package entity

import (
	"time"
)

type Student struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	StudentNo      string     `gorm:"uniqueIndex;size:50;not null" json:"student_no"`
	Name           string     `gorm:"size:50;not null" json:"name"`
	Gender         string     `gorm:"type:ENUM('male','female');" json:"gender"`
	Ethnicity      string     `gorm:"size:50" json:"ethnicity"`
	BirthDate      *time.Time `gorm:"type:date" json:"birth_date"`
	Education      string     `gorm:"size:50" json:"education"`
	Phone          string     `gorm:"size:20" json:"phone"`
	IDCard         string     `gorm:"size:20" json:"id_card"`
	MajorClass     string     `gorm:"size:100" json:"major_class"`
	Type           string     `gorm:"type:ENUM('masses','activist','candidate','probationary','formal');default:'masses'" json:"type"`
	AdmissionDate  *time.Time `gorm:"type:date" json:"admission_date"`
	ConversionDate *time.Time `gorm:"type:date" json:"conversion_date"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
