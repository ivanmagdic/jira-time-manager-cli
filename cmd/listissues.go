package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

type IssueList struct {
	Issues []Issue
}

type Issue struct {
	Id     string
	Fields Fields
}

type Fields struct {
	Summary string
}

var listissuesCmd = &cobra.Command{
	Use:   "list-issues",
	Short: "Check status of deployed artifacts (web, api or database)",
	Long:  `This command can be used together with web, api or database sub-commands to check status of respective deployments`,
	Run:   fetchIssueList,
}

func init() {
	rootCmd.AddCommand(listissuesCmd)
}

func fetchIssueList(*cobra.Command, []string) {
	url = viper.GetString("endpoint") + "search?jql=" + viper.GetString("issue_list")
	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+viper.GetString("auth"))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)

	var issueList IssueList
	err = json.Unmarshal([]byte(sb), &issueList)
	if err != nil {
		return
	}

	var issueIds []string

	for _, issue := range issueList.Issues {
		issueIds = append(issueIds, issue.Id)

		if issue.Id == viper.GetString("active_issue.id") {
			fmt.Println("\033[32m", issue.Id, " - ", issue.Fields.Summary)
			continue
		}

		fmt.Println("\033[0m", issue.Id, " - ", issue.Fields.Summary)
	}

	viper.Set("issues", issueIds)
	viper.WriteConfig()
}
