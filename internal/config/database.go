package config

import (
	"fmt"
)

const (
	databaseHostEnv     = "APP_DB_HOST"
	databasePortEnv     = "APP_DB_PORT"
	databaseSslModeEnv  = "APP_DB_SSL_MODE"
	databaseNameEnv     = "APP_DB_NAME"
	databaseUserEnv     = "APP_DB_USER"
	databasePasswordEnv = "APP_DB_PASSWORD"
)

func DatabaseHost() string {
	return optionalEnv(databaseHostEnv, "localhost")
}

func DatabaseName() string {
	return requiredEnv(databaseNameEnv)
}

func DatabaseSslMode() string {
	return optionalEnv(databaseSslModeEnv, "disable")
}

func DatabasePort() string {
	return optionalEnv(databasePortEnv, "5432")
}

func DatabaseUser() string {
	return requiredEnv(databaseUserEnv)
}

func DatabasePassword() string {
	return requiredEnv(databasePasswordEnv)
}

func DataSourceName() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=%s",
		DatabaseUser(),
		DatabasePassword(),
		DatabaseHost(),
		DatabasePort(),
		DatabaseName(),
		DatabaseSslMode(),
		AppTimezone(),
	)
}
