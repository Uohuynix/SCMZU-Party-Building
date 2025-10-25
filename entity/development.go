package entity

import (
	"time"
)

type Development struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	StudentID        uint       `json:"student_id"`
	ApplyDate        *time.Time `gorm:"type:date" json:"apply_date"`
	ActivistDate     *time.Time `gorm:"type:date" json:"activist_date"`
	CandidateDate    *time.Time `gorm:"type:date" json:"candidate_date"`
	ProbationDate    *time.Time `gorm:"type:date" json:"probation_date"`
	ConversionDate   *time.Time `gorm:"type:date" json:"conversion_date"`
	TransferDate     *time.Time `gorm:"type:date" json:"transfer_date"`
	IntroductionDate *time.Time `gorm:"type:date" json:"introduction_date"`
	Status           string     `gorm:"type:ENUM('masses','activist','candidate','probationary','formal');default:'masses'" json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	Student          Student    `gorm:"foreignKey:StudentID" json:"student,omitempty"`
}
