package configs

import "github.com/muktiarafi/myriadcode-backend/internal/logs"

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

type AppConfig struct {
	*logs.Logger
	WithMigration bool
}
