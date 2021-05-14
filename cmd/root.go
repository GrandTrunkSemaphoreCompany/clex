// CLI commands for running a clex server
package cmd

import (
	"fmt"
	"github.com/GrandTrunkSemaphoreCompany/clex/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gocv.io/x/gocv"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// CameraConfig represents an input camera
type CameraConfig struct {
	Id int
	Url string
	Position int
}

// Config and an initialised gocv.VideoCapture
type VideoInput struct {
	CameraConfig CameraConfig
	VideoCapture *gocv.VideoCapture
}

var (
	id     int
	sink   []string
	source []string
	c      config.Config
	cameras []CameraConfig


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

	viper.UnmarshalKey("cameras", &cameras)

}

func (v CameraConfig) InitVideoInput() (VideoInput, error) {
	var videoInput VideoInput

	url, err := url.Parse(v.Url)
	if err != nil {
		return videoInput, fmt.Errorf("error parsing %v", v.Url)
	}

	if url.Scheme != "usb" {
		return videoInput, fmt.Errorf("Scheme %s not supported", url.Scheme)
	}

	//if url.Scheme == "usb" {
	idPart := strings.Replace(url.Path, "/dev/video", "", -1)
	//fmt.Printf("idPart: %s\n", idPart)
	id, err := strconv.Atoi(idPart)
	if err != nil {
		fmt.Println(err)
		return videoInput, fmt.Errorf("ID not valid %s", idPart)
	}

	videoInput.CameraConfig = v
	capture, err := gocv.OpenVideoCapture(id)
	if err != nil {
		return videoInput, err
	}

	videoInput.VideoCapture = capture
	return videoInput, nil
}
