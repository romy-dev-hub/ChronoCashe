package cashe

import (
	"chronocashe/internal/models"
	"sync"
	"time"
)

type Cache struct {
	data  map[string]models.CasheEntry
	mutex sync.RWMutex
}

// NewCashe creates and returns a new cashe instance
func NewCashe() *Cache {
	return &Cache{
		data: make(map[string]models.CasheEntry),
	}
}

// set stores a new key in the cashe
func (c *Cache) Set(entry models.CasheEntry) {
	c.mutex.Lock()
	defer c.mutex.RUnlock()
	c.data[entry.Key] = entry
}

// delete removes a key from the cashe
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, key)
}

// PruneExpired removes all keys that are completely expired
func (c *Cache) PruneExpired() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for key, entry := range c.data {
		if now.After(entry.AvailableUntil) {
			delete(c.data, key)
		}
	}
}

// GetAllActive returns all currently keys that are active
func (c *Cache) GetAllActive() []models.CasheEntry {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	now := time.Now()
	var active []models.CasheEntry
	for _, entry := range c.data {
		if now.After(entry.AvailableFrom) && now.Before(entry.AvailableUntil) {
			active = append(active, entry)
		}
	}

	return active
}

// Get retrieves a key only if it's within the valid time window
func (c *Cache) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, exists := c.data[key]
	if !exists {
		return "", false
	}

	now := time.Now()
	if now.After(entry.AvailableFrom) && now.Before(entry.AvailableUntil) {
		return entry.Value, true
	}

	return "", false
}
