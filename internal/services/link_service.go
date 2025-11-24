//à remplacer


package services

import (
	"errors"
	"github.com/axellelanca/urlshortener/internal/models"
	"github.com/axellelanca/urlshortener/internal/repository"
)

type LinkService struct {
	LinkRepo  repository.LinkRepository
	ClickRepo repository.ClickRepository // Optionnel selon ton implémentation
}

// NewLinkService crée le service
func NewLinkService(linkRepo repository.LinkRepository, clickRepo repository.ClickRepository) *LinkService {
	return &LinkService{
		LinkRepo:  linkRepo,
		ClickRepo: clickRepo,
	}
}

// CreateShortLink simule la logique métier
func (s *LinkService) CreateShortLink(originalURL string) (*models.Link, error) {
	if originalURL == "" {
		return nil, errors.New("URL vide")
	}
	// Simulation : on renvoie un objet factice
	return &models.Link{
		ShortCode: "XYZ123",
		LongURL:   originalURL,
	}, nil
}

// GetOriginalURL récupère l'URL (Simulation)
func (s *LinkService) GetOriginalURL(code string) (string, error) {
	return "http://google.com", nil
}

// GetLinkStats récupère les stats (Simulation)
func (s *LinkService) GetLinkStats(code string) (*models.Link, error) {
	return &models.Link{
		ShortCode: code,
		ClickCount: 42,
	}, nil
}