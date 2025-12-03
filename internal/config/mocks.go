package config

import (
	"log/slog"

	"github.com/stretchr/testify/mock"
)

type MockConfig struct {
	mock.Mock
}

func (m *MockConfig) GetConfigVars() *ConfigVars {
	args := m.Called()
	return args.Get(0).(*ConfigVars)
}

func (m *MockConfig) GetLogger() *slog.Logger {
	args := m.Called()
	return args.Get(0).(*slog.Logger)
}
