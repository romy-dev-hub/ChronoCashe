package scheduler

import (
	"chronocashe/internal/cache"
	"time"
)

// Start begins periodic cleanup of expired keys
func Start(c *cashe.Cache, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.PruneExpired()
	}
}
