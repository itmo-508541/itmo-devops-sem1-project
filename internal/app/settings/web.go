package settings

import (
	"fmt"
	"project_sem/internal/config"
)

const (
	WebHostDefault = "0.0.0.0"
	WebPortDefault = "8080"

	webHostEnv = "APP_HOST"
	webPortEnv = "APP_PORT"
)

type WebSettings struct {
	Host string
	Port string
}

func NewWebSettings() *WebSettings {
	return &WebSettings{
		Host: config.OptionalEnv(webHostEnv, WebHostDefault),
		Port: config.OptionalEnv(webPortEnv, WebPortDefault),
	}
}

func (c WebSettings) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
