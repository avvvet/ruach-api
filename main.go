package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/avvvet/ruach-api/config"
	"github.com/avvvet/ruach-api/handler"
	"github.com/avvvet/ruach-api/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func main() {
	// load config
	cfg := config.Load()

	// ensure db directory exists
	os.MkdirAll(filepath.Dir(cfg.DBPath), 0755)

	// init db
	store.Init(cfg.DBPath)
	defer store.Close()

	// router
	r := chi.NewRouter()

	// middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.CleanPath)

	// cors — allow Svelte dev and production
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173", // svelte dev
			"http://localhost:4173", // svelte preview
			"http://localhost:3000",
		},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
		MaxAge:         300,
	}))

	// global rate limit — protect GPU
	r.Use(middleware.ThrottleBacklog(10, 20, 30*time.Second))

	// per IP rate limit — 5 requests per minute
	r.Use(httprate.LimitByIP(5, 1*time.Minute))

	// routes
	r.Post("/api/transcribe", handler.Transcribe(cfg))
	r.Get("/api/recent", handler.Recent)
	r.Get("/api/health", handler.Health(nil))

	log.Printf("ruach-api: starting on port %s", cfg.APIPort)
	log.Printf("ruach-api: forwarding to %s", cfg.RuachURL)

	if err := http.ListenAndServe(":"+cfg.APIPort, r); err != nil {
		log.Fatalf("ruach-api: server error: %v", err)
	}
}
