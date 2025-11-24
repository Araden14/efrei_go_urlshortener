package monitor

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/axellelanca/urlshortener/internal/repository"
)

// UrlMonitor gère la surveillance périodique des URLs longues.
type UrlMonitor struct {
	linkRepo    repository.LinkRepository // On utilise l'interface, c'est plus propre
	interval    time.Duration
	knownStates map[uint]bool
	mu          sync.Mutex
}

// NewUrlMonitor crée et retourne une nouvelle instance de UrlMonitor.
func NewUrlMonitor(linkRepo repository.LinkRepository, interval time.Duration) *UrlMonitor {
	return &UrlMonitor{
		linkRepo:    linkRepo,
		interval:    interval,
		knownStates: make(map[uint]bool),
	}
}

// Start lance la boucle de surveillance.
func (m *UrlMonitor) Start() {
	log.Printf("[MONITOR] Démarrage avec un intervalle de %v...", m.interval)
	ticker := time.NewTicker(m.interval)
	defer ticker.Stop()

	// Première vérification immédiate
	m.checkUrls()

	for range ticker.C {
		m.checkUrls()
	}
}

// checkUrls effectue une vérification de l'état.
func (m *UrlMonitor) checkUrls() {
	log.Println("[MONITOR] Vérification des URLs en cours...")

	links, err := m.linkRepo.GetAllLinks()
	if err != nil {
		log.Printf("[MONITOR] ERREUR récupération liens : %v", err)
		return
	}

	for _, link := range links {
		currentState := m.isUrlAccessible(link.LongURL)

		m.mu.Lock()
		previousState, exists := m.knownStates[link.ID]
		m.knownStates[link.ID] = currentState
		m.mu.Unlock()

		if !exists {
			// Première fois qu'on voit ce lien
			continue
		}

		if previousState != currentState {
			// Notification de changement d'état
			log.Printf("[NOTIFICATION] Le lien %s (%s) est passé de %s à %s !",
				link.LongURL, link.ShortCode, formatState(previousState), formatState(currentState))
		}
	}
}

func (m *UrlMonitor) isUrlAccessible(url string) bool {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Head(url)
	if err != nil {
		// log.Printf("[MONITOR] Erreur accès '%s': %v", url, err) // Optionnel pour ne pas polluer les logs
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode >= 200 && resp.StatusCode < 400
}

func formatState(accessible bool) string {
	if accessible {
		return "ACCESSIBLE"
	}
	return "INACCESSIBLE"
}