package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encodeCmd)
}

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encodes a Clacks message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encode: not implemented")
	},
}