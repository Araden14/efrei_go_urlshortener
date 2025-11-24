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
	for event := range clickEventsChan { // Boucle qui lit les événements du channel
		click := models.Click{
			LinkID: event.LinkID,
			Timestamp: event.Timestamp,
			UserAgent: event.UserAgent,
			IPAddress: event.IP,
		}

		// Implémentez ici une gestion d'erreur simple : loggez l'erreur si la persistance échoue.
		// Pour un système en production, une logique de retry
		err := clickRepo.CreateClick(&click)

		err := clickRepo.CreateClick(&click)
		if err != nil {
			// Si une erreur se produit lors de l'enregistrement, logguez-la.
			// L'événement est "perdu" pour ce TP, mais dans un vrai système,
			// vous pourriez le remettre dans une file de retry ou une file d'erreurs.
			log.Printf("ERROR: Failed to save click for LinkID %d (UserAgent: %s, IP: %s): %v",
				event.LinkID, event.UserAgent, event.IP, err)

		} else {
			log.Printf("Click recorded for LinkID %d", event.LinkID)
		}
	}
}