package handlers

import (
	"todo/domain"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server struct
type Server struct {
	domain *domain.Domain
}

func setupMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Timeout(60 * time.Second))
}

// NewServer Server type
func NewServer(domain *domain.Domain) *Server {
	return &Server{domain: domain}
}

// SetupRouter chi
func SetupRouter(domain *domain.Domain) *chi.Mux {
	server := NewServer(domain)

	r := chi.NewRouter()

	setupMiddlewares(r)

	server.setupEndpoints(r)

	return r
}
