package general

import (
	"context"
	"os/signal"
	"project_sem/internal/config"
	"syscall"

	"github.com/sarulabs/di"
)

const (
	ConfigServiceName  = "general:config"
	ContextServiceName = "general:context"

	TimezoneDefault = "Europe/Moscow"

	timezoneEnv = "APP_TIMEZONE"
)

var Services = []di.Def{
	{
		Name:  ConfigServiceName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := &Config{
				Timezone: config.OptionalEnv(timezoneEnv, TimezoneDefault),
			}

			return cfg, nil
		},
	},
	func() di.Def {
		rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

		return di.Def{
			Name:  ContextServiceName,
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return rootCtx, nil
			},
			Close: func(obj interface{}) error {
				stop()
				return nil
			},
		}
	}(),
}
