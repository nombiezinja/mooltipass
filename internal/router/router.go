package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/nombiezinja/mooltipass-internal/internal/config"
)

func NewRouter(cfg config.ConfigInterface) *chi.Mux {
	r := chi.NewRouter()

	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))

	// TODO mount middlewares.

	return r
}
