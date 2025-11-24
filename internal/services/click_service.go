package services

import (
	"github.com/axellelanca/urlshortener/internal/models"
	"github.com/axellelanca/urlshortener/internal/repository"
)

// 👇 VOICI CE QUI MANQUAIT : La définition de la structure
type ClickService struct {
	repo repository.ClickRepository
}

// NewClickService crée une nouvelle instance du service
func NewClickService(repo repository.ClickRepository) *ClickService {
	return &ClickService{
		repo: repo,
	}
}

// RecordClick enregistre un clic (Bouchon/Wrapper vers le repo)
func (s *ClickService) RecordClick(click *models.Click) error {
	// On appelle la méthode du repository (assure-toi que CreateClick existe dans ton interface repository)
	// Si ton repository s'appelle SaveClick, change ici.
	return s.repo.CreateClick(click)
}

// GetClicksCount compte les clics pour un lien
func (s *ClickService) GetClicksCount(linkID uint) (int, error) {
	return s.repo.CountClicksByLinkID(linkID)
}