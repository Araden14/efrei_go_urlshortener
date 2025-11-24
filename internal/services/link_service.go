package services

import (
	"crypto/rand"

	"github.com/axellelanca/urlshortener/internal/models"
	"github.com/axellelanca/urlshortener/internal/repository"
)

// Définition du jeu de caractères pour la génération des codes courts.
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// LinkService est une structure qui g fournit des méthodes pour la logique métier des liens.
// Elle détient linkRepo qui est une référence vers une interface LinkRepository.
// IMPORTANT : Le champ doit être du type de l'interface (non-pointeur).
type LinkService struct {
	LinkRepo repository.LinkRepository
}

// NewLinkService crée et retourne une nouvelle instance de LinkService.
func NewLinkService(linkRepo repository.LinkRepository) *LinkService {
	return &LinkService{
		LinkRepo: linkRepo,
	}
}

// GenerateShortCode est une méthode rattachée à LinkService
// Elle génère un code court aléatoire d'une longueur spécifiée. Elle prend une longueur en paramètre et retourne une string et une erreur
// Il utilise le package 'crypto/rand' pour éviter la prévisibilité.
// Je vous laisse chercher un peu :) C'est faisable en une petite dizaine de ligne
func GenerateShortCode(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convertir les bytes aléatoires en caractères du charset
	for i := 0; i < length; i++ {
		bytes[i] = charset[int(bytes[i])%len(charset)]
	}

	return string(bytes), nil
}

// CreateLink crée un nouveau lien raccourci.
// Il génère un code court unique, puis persiste le lien dans la base de données.
func (s *LinkService) CreateLink(longURL string) (*models.Link, error) {
	// TODO 1: Implémenter la logique de retry pour générer un code court unique.
	// Essayez de générer un code, vérifiez s'il existe déjà en base, et retentez si une collision est trouvée.
	// Limitez le nombre de tentatives pour éviter une boucle infinie.

	// TODO Créer une variable shortcode pour stocker le shortcode créé
	var shortcode string

	// TODO Définir un nombre maximum (5) de tentative pour trouver un code unique  (maxRetries)
	maxRetries := 5
	var err error

	// Boucle de retry pour trouver un code unique
	for i := 0; i < maxRetries; i++ {
		// Générer un code court de 6 caractères
		shortcode, err = GenerateShortCode(6)
		if err != nil {
			return nil, err
		}

		// Vérifier si ce code existe déjà
		existingLink, err := s.LinkRepo.GetLinkByShortCode(shortcode)
		if err != nil {
			// Si l'erreur est "record not found", c'est parfait - le code est unique
			// On suppose que GORM renvoie gorm.ErrRecordNotFound
			if existingLink == nil {
				// Code unique trouvé !
				break
			}
		}

		// Si on arrive ici et que c'est le dernier essai, on échoue
		if i == maxRetries-1 {
			return nil, err
		}
	}

	// Créer le modèle Link
	link := &models.Link{
		Shortcode: shortcode,
		LongURL:   longURL,
	}

	// Persister dans la base de données
	err = s.LinkRepo.CreateLink(link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

// GetLinkByShortCode délègue au repository.
func (s *LinkService) GetLinkByShortCode(shortCode string) (*models.Link, error) {
	return s.LinkRepo.GetLinkByShortCode(shortCode)
}

// GetClicksCount (Ajouté car ton Handler l'utilise)
func (s *LinkService) GetClicksCount(linkID uint) (int, error) {
	return s.LinkRepo.CountClicksByLinkID(linkID)
}

// GetLinkStats combine les infos du lien et des clics (Bonus si besoin)
func (s *LinkService) GetLinkStats(shortCode string) (*models.Link, int, error) {
	// 1. Récupérer le lien
	link, err := s.LinkRepo.GetLinkByShortCode(shortCode)
	if err != nil {
		return nil, 0, err
	}

	// 2. Compter les clics
	count, err := s.LinkRepo.CountClicksByLinkID(link.ID)
	if err != nil {
		return nil, 0, err
	}

	return link, count, nil
}
