package settings

import (
	"fmt"
	"project_sem/internal/config"
)

const (
	DatabaseHostDefault    = "localhost"
	DatabasePortDefault    = "5432"
	DatabaseSslModeDefault = "disable"

	databaseHostEnv     = "APP_DB_HOST"
	databasePortEnv     = "APP_DB_PORT"
	databaseSslModeEnv  = "APP_DB_SSL_MODE"
	databaseNameEnv     = "APP_DB_NAME"
	databaseUserEnv     = "APP_DB_USER"
	databasePasswordEnv = "APP_DB_PASSWORD"
)

type DatabaseSettings struct {
	Host     string
	Port     string
	SslMode  string
	Database string
	User     string
	Password string
	Timezone string
}

func NewDatabaseSettings() *DatabaseSettings {
	return &DatabaseSettings{
		Host:     config.OptionalEnv(databaseHostEnv, DatabaseHostDefault),
		Port:     config.OptionalEnv(databasePortEnv, DatabasePortDefault),
		SslMode:  config.OptionalEnv(databaseSslModeEnv, DatabaseSslModeDefault),
		Database: config.RequiredEnv(databaseNameEnv),
		User:     config.RequiredEnv(databaseUserEnv),
		Password: config.RequiredEnv(databasePasswordEnv),
		Timezone: config.OptionalEnv(timezoneEnv, TimezoneDefault),
	}
}

func (c DatabaseSettings) DataSourceName() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.SslMode,
		c.Timezone,
	)
}
