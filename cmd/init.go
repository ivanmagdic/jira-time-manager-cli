package cmd

import (
	"encoding/base64"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var (
	url      string
	username string
	password string
	initCmd  = &cobra.Command{
		Use:   "init",
		Short: "Create basic authentication credentials",
		Long:  `This command can be used together with web, api or database sub-commands to check status of respective deployments`,
		Run:   initializeConfiguration,
	}
)

func init() {
	initCmd.Flags().StringVar(&url, "url", "", "JIRA server URL")
	initCmd.Flags().StringVarP(&username, "username", "u", "", "JIRA username")
	initCmd.Flags().StringVarP(&password, "password", "p", "", "JIRA password")

	initCmd.MarkFlagRequired("url")
	initCmd.MarkFlagRequired("username")
	initCmd.MarkFlagRequired("password")

	rootCmd.AddCommand(initCmd)
}

func initializeConfiguration(cmd *cobra.Command, args []string) {
	viper.Set("baseUrl", url)
	viper.Set("auth", base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
	viper.Set("endpoint", url+"/rest/api/2/")
	viper.Set("issue_list", "assignee=currentuser()%20AND%20resolution%20=%20Unresolved%20order%20by%20updated%20DESC")
	err := viper.WriteConfig()
	if err != nil {
		log.Println("Error:", err.Error())
		return
	}
}
