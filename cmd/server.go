package cmd

import (
	"github.com/GrandTrunkSemaphoreCompany/clex/clacks/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCommand)
}

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Runs the standalone Clex server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start(c)
	},
}