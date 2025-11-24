package models

import "time"

// ClickEvent est utilisé pour passer les données entre le Handler et le Worker via le channel
type ClickEvent struct {
	LinkID    uint
	Timestamp time.Time
	IP        string
	UserAgent string
}