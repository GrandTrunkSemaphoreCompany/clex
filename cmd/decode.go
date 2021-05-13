package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decodeCmd)
}

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes a Clacks message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("decode: not implemented")
	},
}