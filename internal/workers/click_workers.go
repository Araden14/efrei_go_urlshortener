//à remplacer

package workers

import "fmt"

type ClickWorker struct{}

func NewClickWorker() *ClickWorker {
	return &ClickWorker{}
}

func (w *ClickWorker) Start() {
	fmt.Println("👷 Worker (Simulation) : Prêt")
}

func (w *ClickWorker) AddClick(shortCode, ip, userAgent string) {
	// Ne fait rien pour l'instant
}