package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/nombiezinja/mooltipass-internal/internal/config"
	"github.com/stretchr/testify/assert"
)

// TestNewServiceSuccess ensures New returns a valid Service and Server
func TestNewServiceWithMockConfig(t *testing.T) {
	expectedVars := &config.ConfigVars{
		Environment:      "test",
		AppName:          "mooltipass-test",
		JwtSigningSecret: "secret",
		LogLevel:         "info",
		Port:             "8080",
		DatabaseURL:      "db-url-test",
	}
	mockCfg := &config.MockConfig{}
	mockCfg.On("GetConfigVars").Return(expectedVars)

	r := chi.NewMux()
	svc, srv := New(mockCfg, r)
	assert.NotNil(t, svc)
	assert.NotNil(t, srv)
	// Compare the values of ConfigVars
	actualVars := srv.Config.GetConfigVars()
	assert.EqualValues(t, expectedVars, actualVars)
}

// TestHandlerMethodsNotImplemented checks all handler stubs return 501
func TestHandlerMethodsNotImplemented(t *testing.T) {
	cfg := &config.Config{}
	r := chi.NewMux()
	_, srv := New(cfg, r)

	// Helper to test handler
	testNotImplemented := func(handler func(http.ResponseWriter, *http.Request), name string) {
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		handler(w, req)
		assert.Equal(t, http.StatusNotImplemented, w.Code, name)
	}

	testNotImplemented(srv.HealthCheck, "HealthCheck")
	testNotImplemented(srv.GetFavicon, "GetFavicon")
	testNotImplemented(srv.Register, "Register")
	testNotImplemented(srv.Login, "Login")
	testNotImplemented(srv.ListPrincipals, "ListPrincipals")
	testNotImplemented(srv.CreatePrincipal, "CreatePrincipal")
	testNotImplemented(srv.DeletePrincipal, "DeletePrincipal")
	testNotImplemented(srv.GetPrincipal, "GetPrincipal")
	testNotImplemented(srv.UpdatePrincipal, "UpdatePrincipal")
	testNotImplemented(srv.ListPosts, "ListPosts")
	testNotImplemented(srv.DeletePost, "DeletePost")
	testNotImplemented(srv.GetPost, "GetPost")
}

// TestNewServiceFailure checks New with nil config returns valid objects
func TestNewServiceFailure(t *testing.T) {
	r := chi.NewMux()
	svc, srv := New(nil, r)
	assert.NotNil(t, svc)
	assert.NotNil(t, srv)
	assert.Nil(t, srv.Config)
}
