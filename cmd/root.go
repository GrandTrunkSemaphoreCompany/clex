// CLI commands for running a clex server
package cmd

import (
	"fmt"
	"github.com/GrandTrunkSemaphoreCompany/clex/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	id     int
	sink   []string
	source []string
	c      config.Config

	rootCmd = &cobra.Command{
		Use:   "clex",
		Short: "Clex works with a Clacks system to send messages via visual semaphore",
		Long: `Clex is a computer based application for interfacing with a visual 
semaphore system. The semaphore system is based upon the Clacks
from Terry Pratchett's Discworld novels.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().IntVar(&id, "id", 0, "Numeric ID for this clex server")
	rootCmd.PersistentFlags().StringArrayVarP(&sink, "sink", "s", []string{}, "sink(s) to configure")
	rootCmd.PersistentFlags().StringArrayVarP(&source, "source", "c", []string{}, "source(s) to configure")

	viper.BindPFlag("id", rootCmd.PersistentFlags().Lookup("id"))
	viper.BindPFlag("sink", rootCmd.PersistentFlags().Lookup("sink"))
	viper.BindPFlag("source", rootCmd.PersistentFlags().Lookup("source"))

	//rootCmd.PersistentFlags().StringVar(&id, "id", 0, "Numeric ID for this clex server")
	//rootCmd.PersistentFlags().Int(&id, "id", 0, "Numeric ID for this clex server")
	//flag.IntVar(&id, "flagname", 1234, "help message for flagname")
}

func initConfig() {
	viper.SetConfigName("config")      // name of config file (without extension)
	viper.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/clex/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.clex") // call multiple times to add many search paths
	viper.AddConfigPath(".")           // optionally look for config in the working directory
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
