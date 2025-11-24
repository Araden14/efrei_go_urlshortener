package repository

import (
	// 👇 Import correct des modèles
	"github.com/axellelanca/urlshortener/internal/models"
	"gorm.io/gorm"
)

type ClickRepository struct {
	DB *gorm.DB
}

// Save : Sauvegarde un clic (Simulation)
func (r *ClickRepository) Save(click *models.Click) error {
	return nil
}