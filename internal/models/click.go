package models

import (
	"time"
	"gorm.io/gorm"
)

// Click : Définition de ce qu'est un clic
type Click struct {
	gorm.Model
	LinkID    uint      `json:"link_id"`
	UserAgent string    `json:"user_agent"`
	IPAddress string    `json:"ip_address"`
	ClickedAt time.Time `json:"clicked_at"`
}