package monitor

import (
	"fmt"
	"time"
	"github.com/axellelanca/urlshortener/internal/services"
)

type URLMonitor struct {
	Service *services.LinkService
}

func NewURLMonitor(service *services.LinkService) *URLMonitor {
	return &URLMonitor{Service: service}
}

func (m *URLMonitor) Start() {
	fmt.Println("👀 Moniteur (Simulation) : Prêt")
	// Boucle infinie non bloquante
	go func() {
		for {
			time.Sleep(10 * time.Minute)
		}
	}()
}