package services

import (
	"github.com/sarulabs/di"
	"github.com/spf13/cobra"
)

const (
	CommandRootServiceName = "application:command.root"
)

var CommandServices = []di.Def{
	{
		Name:  CommandRootServiceName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			rootCmd := &cobra.Command{
				Short: "Final project 1st semester (Andrey Mindubaev, id:508541)",
			}
			rootCmd.AddCommand(ctn.Get(MigrateCommandServiceName).(*cobra.Command))
			rootCmd.AddCommand(ctn.Get(StartServerCommandServiceName).(*cobra.Command))

			return rootCmd, nil
		},
	},
}
