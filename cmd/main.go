package main

import (
	"fmt"
	"log"
	"project_sem/internal/app/command"

	"github.com/spf13/cobra"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	rootCmd := &cobra.Command{
		Short: "Final project 1st semester (Andrey Mindubaev, id:508541)",
	}
	rootCmd.AddCommand(command.NewMigrate())
	rootCmd.AddCommand(command.NewStartServer())

	if err := rootCmd.Execute(); err != nil {
		log.Println(fmt.Errorf("rootCmd.Execute: %w", err))
	}
}
