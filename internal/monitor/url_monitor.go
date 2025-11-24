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
	linkRepo    repository.GormLinkRepository // Pour récupérer les URLs à surveiller
	interval    time.Duration             // Intervalle entre chaque vérification (ex: 5 minutes)
	knownStates map[uint]bool             // État connu de chaque URL: map[LinkID]estAccessible (true/false)
	mu          sync.Mutex                // Mutex pour protéger l'accès concurrentiel à knownStates
}

// NewUrlMonitor crée et retourne une nouvelle instance de UrlMonitor.
func NewUrlMonitor(linkRepo repository.GormLinkRepository, interval time.Duration) *UrlMonitor {
	pUrlMonitor := &UrlMonitor{
		linkRepo: linkRepo,
		interval: interval,
		knownStates: make(map[uint]bool),
	}
	return pUrlMonitor;
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

	// Gérer l'erreur si la récupération échoue.
	// Si erreur : log.Printf("[MONITOR] ERREUR lors de la récupération des liens pour la surveillance : %v", err)
	links, err := m.linkRepo.GetAllLinks();
	if (err != nil) {
		log.Printf("[MONITOR] ERREUR lors de la récupération des liens pour la surveillance : %v", err)
	}

	for _, link := range links {
		currentState := m.isUrlAccessible(link.LongURL);

		m.mu.Lock()
		previousState, exists := m.knownStates[link.ID]
		m.knownStates[link.ID] = currentState
		m.mu.Unlock()

		if !exists {
			log.Printf("[MONITOR] État initial pour le lien %s (%s) : %s",
				link.Shortcode, link.LongURL, formatState(currentState))
			continue
		}

		// Si l'état a changé, générer une fausse notification dans les logs.
		if (previousState != currentState) {
			log.Printf("[NOTIFICATION] Le lien %s (%s) est passé de %s à %s !", link.LongURL, link.Shortcode, previousState, currentState);
		}

	}
}

func (m *UrlMonitor) isUrlAccessible(url string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Un code de statut 2xx ou 3xx indique que l'URL est accessible.
	resp, err := client.Head(url)
	if (err != nil) {
		log.Printf("[MONITOR] Erreur d'accès à l'URL '%s': %v", url, err)
	}

	defer resp.Body.Close()

	// Déterminer l'accessibilité basée sur le code de statut HTTP.
	return resp.StatusCode >= 200 && resp.StatusCode < 400 // Codes 2xx ou 3xx
}

func formatState(accessible bool) string {
	if accessible {
		return "ACCESSIBLE"
	}
	return "INACCESSIBLE"
}