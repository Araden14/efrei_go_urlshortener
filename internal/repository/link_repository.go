package repository

import (
	// 👇 C'est LÀ que ça bloquait. On importe tes modèles correctement.
	"github.com/axellelanca/urlshortener/internal/models"
	"gorm.io/gorm"
)

type LinkRepository struct {
	DB *gorm.DB
}

// Create : Sauvegarde un lien (Simulation)
func (r *LinkRepository) Create(link *models.Link) error {
	return nil
}

// FindByShortCode : Cherche un lien (Simulation)
func (r *LinkRepository) FindByShortCode(code string) (*models.Link, error) {
	// On renvoie un lien vide pour que ça compile
	return &models.Link{}, nil
}

// IncrementClicks : Compte les clics (Simulation)
func (r *LinkRepository) IncrementClicks(link *models.Link) error {
	return nil
}

// GetAll : Récupère tout (Simulation)
func (r *LinkRepository) GetAll() ([]models.Link, error) {
	return []models.Link{}, nil
}