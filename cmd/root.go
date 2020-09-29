package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

var (
	version    string
	configPath string
)

var rootCmd = &cobra.Command{
	SilenceUsage: true,
	Short:        "Start Order Server",

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit
		log.Warn().Msg("Shutting down server...")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "c", "./internal/pkg/config", "Config Path")
}

func Execute() {
	rootCmd.AddCommand(
		grpcCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
