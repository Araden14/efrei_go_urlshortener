package workers

import (
	"log"
	"github.com/axellelanca/urlshortener/internal/models"
	"github.com/axellelanca/urlshortener/internal/repository"
)

// StartClickWorkers lance les workers.
func StartClickWorkers(workerCount int, clickEventsChan <-chan models.ClickEvent, clickRepo repository.ClickRepository) {
	log.Printf("Starting %d click worker(s)...", workerCount)
	for i := 0; i < workerCount; i++ {
		go clickWorker(clickEventsChan, clickRepo)
	}
}

func clickWorker(clickEventsChan <-chan models.ClickEvent, clickRepo repository.ClickRepository) {
	for event := range clickEventsChan {
		click := models.Click{
			LinkID:    event.LinkID,
			ClickedAt: event.Timestamp, // Attention: Vérifie si ton modèle s'appelle Timestamp ou ClickedAt
			UserAgent: event.UserAgent,
			IPAddress: event.IP,        // Attention: Vérifie si ton modèle s'appelle IP ou IPAddress
		}

		err := clickRepo.CreateClick(&click)
		if err != nil {
			log.Printf("ERROR: Failed to save click: %v", err)
		} else {
			log.Printf("Click recorded for LinkID %d", event.LinkID)
		}
	}
}