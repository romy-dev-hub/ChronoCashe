package main

import (
	"chronocashe/internal/api"
	"chronocashe/internal/cache"
	"chronocashe/internal/scheduler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	// initialize in-memory cashe
	c := cashe.NewCashe()

	// start background scheduler to clean expired keys
	go scheduler.Start(c, 30*time.Second)

	// setup api router
	handler := api.NewAPIHandler(c)
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Route("/cashe", func(r chi.Router) {
		r.Get("/", handler.ListKeys)
		r.Get("/{key}", handler.GetKey)
		r.Put("/{key}", handler.SetKey)
		r.Delete("/{key}", handler.DeleteKey)
	})

	// start server
	log.Println("ChronoCashe server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
