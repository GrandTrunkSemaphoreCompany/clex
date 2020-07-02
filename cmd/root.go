// CLI commands for running a clex server
package cmd

import (
	"fmt"
	"github.com/GrandTrunkSemaphoreCompany/clex/clacks"
	"github.com/spf13/cobra"
	"os"
)

var (
	id      int
	shutter []string
	camera  []string

	rootCmd = &cobra.Command{
		Use:   "clex",
		Short: "Clex works with a Clacks system to send messages via visual semaphore",
		Long: `Clex is a computer based application for interfacing with a visual 
semaphore system. The semaphore system is based upon the Clacks
from Terry Pratchett's Discworld novels.`,
		Run: func(cmd *cobra.Command, args []string) {

			// FIXME: Refactor into single lib method
			fmt.Println("shutters:")
			for _, v := range shutter {
				fmt.Println(v)
			}

			fmt.Println("camera:")
			for _, v := range camera {
				fmt.Println(v)
			}

			clacks.Start()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVar(&id, "id", 0, "Numeric ID for this clex server")
	rootCmd.PersistentFlags().StringArrayVarP(&shutter, "shutter", "s", []string{}, "Shutter(s) to configure")
	rootCmd.PersistentFlags().StringArrayVarP(&camera, "camera", "c", []string{}, "Camera(s) to configure")
	//rootCmd.PersistentFlags().StringVar(&id, "id", 0, "Numeric ID for this clex server")
	//rootCmd.PersistentFlags().Int(&id, "id", 0, "Numeric ID for this clex server")
	//flag.IntVar(&id, "flagname", 1234, "help message for flagname")
}
