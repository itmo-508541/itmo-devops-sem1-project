package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

const (
	dotEnv      = ".env"
	dotEnvLocal = ".env.local"
)

var loadOnce sync.Once

func loadEnv() error {
	if err := godotenv.Overload(dotEnv); err != nil {
		return err
	}

	if _, err := os.Stat(dotEnvLocal); err == nil {
		if err := godotenv.Overload(dotEnvLocal); err != nil {
			return err
		}
	}

	return nil
}

func lookupEnv(key string) (string, bool) {
	loadOnce.Do(func() {
		err := loadEnv()
		if err != nil {
			panic(err)
		}
	})

	return os.LookupEnv(key)
}

func optionalEnv(key string, defaultValue string) string {
	value, ok := lookupEnv(key)
	if !ok {
		value = defaultValue
	}

	return value
}

func requiredEnv(key string) string {
	value, ok := lookupEnv(key)
	if !ok {
		panic(fmt.Errorf("env.%s is required", key))
	}

	return value
}
