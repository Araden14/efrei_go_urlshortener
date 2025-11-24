package models

import (
	"gorm.io/gorm"
)

// Link : Définition de ce qu'est un lien
type Link struct {
	gorm.Model
	ShortCode  string `gorm:"uniqueIndex;not null" json:"short_code"`
	LongURL    string `gorm:"not null" json:"long_url"`
	ClickCount int64  `gorm:"default:0" json:"click_count"`
	IsActive   bool   `gorm:"default:true" json:"is_active"`
}