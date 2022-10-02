package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "jtm",
		Short: "jtm: A simple JIRA time manager CLI application",
		Long:  `jtm: A JIRA time manager CLI application to list, start and stop your issues`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".jira-time-manager")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SafeWriteConfig()
			log.Println("New config file created")
		} else {
			// Config file was found but another error was produced
			log.Println("Error while creating config file:", err.Error())
		}
	}
}
