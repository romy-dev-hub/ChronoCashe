package api

import (
	"chronocashe/internal/cache"
	"chronocashe/internal/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

// APIHandler struct
type APIHandler struct {
	Cache *cashe.Cache
}

// NewAPIHandler creates a new handler with injected cashe
func NewAPIHandler(c *cashe.Cache) *APIHandler {
	return &APIHandler{Cache: c}
}

// PUT /cashe/{key}
func (h *APIHandler) SetKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	var payload struct {
		Value          string `json:"value"`
		AvailableFrom  string `json:"available_from"`
		AvailableUntil string `json:"available_until"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	fromTime, err1 := time.Parse(time.RFC3339, payload.AvailableFrom)
	untilTime, err2 := time.Parse(time.RFC3339, payload.AvailableUntil)
	if err1 != nil || err2 != nil || !fromTime.Before(untilTime) {
		http.Error(w, "Invalid time window", http.StatusBadRequest)
		return
	}

	entry := models.CasheEntry{
		Key:            key,
		Value:          payload.Value,
		AvailableFrom:  fromTime,
		AvailableUntil: untilTime,
	}
	h.Cache.Set(entry)
	w.WriteHeader(http.StatusCreated)

}

// GET /cashe/{key}
func (h *APIHandler) GetKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	value, ok := h.Cache.Get(key)
	if !ok {
		http.Error(w, "key not available or expired", http.StatusNotFound)
		return
	}

	resp := map[string]string{"key": key, "value": value}
	json.NewEncoder(w).Encode(resp)
}

// DELETE /cashe/{key}
func (h *APIHandler) DeleteKey(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	h.Cache.Delete(key)
	w.WriteHeader(http.StatusNoContent)
}

// GET /cashe
func (h *APIHandler) ListKeys(w http.ResponseWriter, r *http.Request) {
	h.Cache.PruneExpired()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Cache.GetAllActive())
}
