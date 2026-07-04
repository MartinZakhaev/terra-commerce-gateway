package httpserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MartinZakhaev/terra-commerce-gateway/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	config config.Config
	http   *http.Server
}

func New(cfg config.Config) *Server {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(cfg.RequestTimeout))

	router.Get("/health/live", healthHandler("live"))
	router.Get("/health/ready", healthHandler("ready"))

	return &Server{
		config: cfg,
		http: &http.Server{
			Addr:              cfg.ListenAddress,
			Handler:           router,
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       15 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
	}
}

func (s *Server) HTTP() *http.Server {
	return s.http
}

func healthHandler(status string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": status})
	}
}
