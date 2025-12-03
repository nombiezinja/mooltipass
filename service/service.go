package service

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nombiezinja/mooltipass-internal/internal/config"
	"github.com/nombiezinja/mooltipass-internal/internal/generatedserver"
)

type Service struct {
	Handler *http.Handler
}

// New returns a new service, inject config and mux as dependencies
// to enable testing.
func New(cfg config.ConfigInterface, r *chi.Mux) (*Service, Server) {
	s := Server{
		Config: cfg,
	}

	opts := generatedserver.ChiServerOptions{
		BaseRouter: r,
	}

	handler := generatedserver.HandlerWithOptions(s, opts)

	return &Service{
		Handler: &handler,
	}, s
}

type Server struct {
	Config config.ConfigInterface
}

// --- System ---
func (s Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) GetFavicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// --- Auth ---
func (s Server) Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) RefreshFacebookToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) VerifyFacebookToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// --- Principals ---
func (s Server) ListPrincipals(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) CreatePrincipal(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) DeletePrincipal(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) GetPrincipal(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) UpdatePrincipal(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// --- Posts ---
func (s Server) ListPosts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) GetPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
