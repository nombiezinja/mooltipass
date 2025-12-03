package config

type ConfigVars struct {
	AppName          string
	Environment      string
	JwtSigningSecret string
	LogLevel         string
	Port             string
	DatabaseURL      string
}
