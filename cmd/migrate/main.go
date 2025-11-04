package main

import (
	"fmt"
	"log"

	"project_sem/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			log.Fatal(panicErr)
		}
	}()

	m, err := migrate.New("file://migrations", config.DataSourceName())
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		fmt.Println(err)
	} else {
		version, _, _ := m.Version()
		fmt.Println("Migrated to version #", version)
	}
}
