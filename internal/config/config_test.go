package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustLoadConfigVarsSuccess(t *testing.T) {
	t.Setenv("APP_NAME", "yeet-test")
	t.Setenv("LOG_LEVEL", "info")
	t.Setenv("DATABASE_URL", "db-url-test")
	t.Setenv("ENVIRONMENT", "test")
	t.Setenv("PORT", "3000")
	t.Setenv("JWT_SIGNING_SECRET", "jwt-signing-secret-test")

	result := MustLoadConfigVars()
	expected := &ConfigVars{
		AppName:          "yeet-test",
		DatabaseURL:      "db-url-test",
		Environment:      "test",
		JwtSigningSecret: "jwt-signing-secret-test",
		LogLevel:         "info",
		Port:             "3000",
	}
	assert.EqualValues(t, expected, result)
}

func TestMustLoadConfigVarsFailure(t *testing.T) {
	t.Setenv("APP_NAME", "yeet-test")
	t.Setenv("LOG_LEVEL", "info")
	t.Setenv("ENVIRONMENT", "test")
	t.Setenv("PORT", "3000")

	assert.Panics(t, func() { MustLoadConfigVars() }, "TestMustLoadConfigVarsFailure - Missing env vars should result in panic but did not")
}
