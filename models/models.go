package models

import "gorm.io/gorm"

// Fact represents a trivia question and its answer
type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"type:text;not null;default:null"`
	Answer string `json:"answer" gorm:"type:text;not null;default:null"`
}