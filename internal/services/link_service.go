package services

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"

	"gorm.io/gorm"

	"github.com/axellelanca/urlshortener/internal/models"
	"github.com/axellelanca/urlshortener/internal/repository"
)

// Définition du jeu de caractères pour la génération des codes courts.
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// LinkService gère la logique métier des liens.
type LinkService struct {
	// Attention à la majuscule ici pour correspondre au NewLinkService
	linkRepo repository.LinkRepository
}

// NewLinkService crée et retourne une nouvelle instance de LinkService.
func NewLinkService(linkRepo repository.LinkRepository) *LinkService {
	return &LinkService{
		linkRepo: linkRepo,
	}
}

// GenerateShortCode génère un code aléatoire sécurisé.
func GenerateShortCode(length int) (string, error) {
	// On crée un tableau de bytes de la bonne longueur
	b := make([]byte, length)
	
	// Pour chaque case, on tire un caractère au hasard
	for i := range b {
		// rand.Int retourne un nombre sécurisé entre 0 et len(charset)
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		// On prend le caractère correspondant à l'index tiré
		b[i] = charset[num.Int64()]
	}
	return string(b), nil
}

// CreateShortLink (Renommé pour correspondre à ton Handler)
func (s *LinkService) CreateShortLink(longURL string) (*models.Link, error) {
	var shortCode string
	maxRetries := 5
    // (La ligne "var err error" a disparu ici)

	// 1. Boucle de retry pour garantir l'unicité
	for i := 0; i < maxRetries; i++ {
        // ...
		// Génère un code de 6 caractères
		code, errGen := GenerateShortCode(6)
		if errGen != nil {
			return nil, fmt.Errorf("erreur génération code: %v", errGen)
		}

		// Vérifie si le code existe déjà
		_, errRepo := s.linkRepo.GetLinkByShortCode(code)

		if errRepo != nil {
			// Si l'erreur est "Record Not Found", C'EST UNE BONNE NOUVELLE !
			// Ça veut dire que le code est libre.
			if errors.Is(errRepo, gorm.ErrRecordNotFound) {
				shortCode = code
				break // On sort de la boucle, on a trouvé notre code unique
			}
			// Vraie erreur de base de données
			return nil, fmt.Errorf("erreur BDD: %w", errRepo)
		}

		// Si on arrive ici, c'est que errRepo == nil, donc le lien existe déjà (collision).
		log.Printf("⚠️ Collision détectée pour %s, nouvelle tentative (%d/%d)...", code, i+1, maxRetries)
	}

	// Si après 5 essais on a toujours rien
	if shortCode == "" {
		return nil, errors.New("impossible de générer un code unique après plusieurs tentatives")
	}

	// 2. Création et sauvegarde
	link := &models.Link{
		LongURL:   longURL,
		ShortCode: shortCode,
		IsActive:  true,
	}

	// Appel au repository pour sauvegarder
	if err := s.linkRepo.CreateLink(link); err != nil {
		return nil, err
	}

	return link, nil
}

// GetLinkByShortCode délègue au repository.
func (s *LinkService) GetLinkByShortCode(shortCode string) (*models.Link, error) {
	return s.linkRepo.GetLinkByShortCode(shortCode)
}

// GetClicksCount (Ajouté car ton Handler l'utilise)
func (s *LinkService) GetClicksCount(linkID uint) (int, error) {
	return s.linkRepo.CountClicksByLinkID(linkID)
}

// GetLinkStats combine les infos du lien et des clics (Bonus si besoin)
func (s *LinkService) GetLinkStats(shortCode string) (*models.Link, int, error) {
	// 1. Récupérer le lien
	link, err := s.linkRepo.GetLinkByShortCode(shortCode)
	if err != nil {
		return nil, 0, err
	}

	// 2. Compter les clics
	count, err := s.linkRepo.CountClicksByLinkID(link.ID)
	if err != nil {
		return nil, 0, err
	}

	return link, count, nil
}