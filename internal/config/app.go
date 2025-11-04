package config

const (
	AppPortEnv     = "APP_PORT"
	AppTimezoneEnv = "APP_TIMEZONE"
)

func AppPort() string {
	return optionalEnv(AppPortEnv, "8080")
}

func AppTimezone() string {
	return optionalEnv(AppTimezoneEnv, "Europe/Moscow")
}
