package main

import (
	"fmt"
	"log"
	"project_sem/internal/services/application"
	"project_sem/internal/services/database"
	"project_sem/internal/services/general"
	"project_sem/internal/services/web"

	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
)

func main() {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			log.Fatal(panicErr)
		}
	}()

	builder, err := di.NewBuilder(di.App)
	if err != nil {
		log.Fatal(fmt.Errorf("di.NewBuilder: %w", err))
	}
	for _, services := range [][]di.Def{general.Services, application.Services, database.Services, web.Services} {
		if err := builder.Add(services...); err != nil {
			log.Fatal(fmt.Errorf("builder.Add: %w", err))
		}
	}

	ctn := builder.Build()
	defer ctn.DeleteWithSubContainers()

	rootCmd := ctn.Get(application.CommandRootServiceName).(*cobra.Command)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(fmt.Errorf("rootCmd.Execute: %w", err))
	}
}
