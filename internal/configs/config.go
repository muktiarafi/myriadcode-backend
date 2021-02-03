package configs

import "log"

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

type AppConfig struct {
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
}
