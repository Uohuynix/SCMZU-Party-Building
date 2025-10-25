package entity

import (
	"time"
)

type Reward struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	StudentID uint       `json:"student_id"`
	Type      string     `gorm:"type:ENUM('reward','punishment');not null" json:"type"`
	Content   string     `gorm:"size:255;not null" json:"content"`
	Date      *time.Time `gorm:"type:date" json:"date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Student   Student    `gorm:"foreignKey:StudentID" json:"student,omitempty"`
}
