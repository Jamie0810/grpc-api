package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/cmd/app"
)

var grpcCmd = &cobra.Command{
	Use:           "grpc",
	Short:         "Start GRPC Server",
	SilenceUsage:  true,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := app.Initialize(configPath)
		if err != nil {
			log.Fatal(err)
		}

		err = app.Launch(version)
		if err != nil {
			log.Fatal(err)
		}
	},
}
