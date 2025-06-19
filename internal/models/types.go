package models

import "time"

// CasheEntry defines the structure of a single cashed key
type CasheEntry struct {
	Key            string    `json:"key"`
	Value          string    `json:"value"`
	AvailableFrom  time.Time `json:"available_from"`
	AvailableUntil time.Time `json:"available_until"`
}
