package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"time"
)

type WorkLog struct {
	timeSpent int64  `json:"timeSpentSeconds"`
	started   string `json:"started"`
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop work time on the active task",
	Run:   stopProgress,
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func stopProgress(cmd *cobra.Command, args []string) {
	var activeIssueId = viper.GetString("active_issue.id")

	if activeIssueId == "" {
		log.Println("There is no active task set")
		return
	}

	url = viper.GetString("endpoint") + "issue/" + activeIssueId + "/worklog"
	var client http.Client

	body, err := json.Marshal(map[string]any{
		"started":          viper.GetString("active_issue.started_at"),
		"timeSpentSeconds": time.Now().Unix() - viper.GetInt64("active_issue.started_at_unix"),
	})

	fmt.Println(url)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Basic "+viper.GetString("auth"))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(respBody)

	log.Println(sb)

	viper.Set("active_issue.id", "")
	viper.Set("active_issue.started_at", "")
	viper.Set("active_issue.started_at_unix", "")
	viper.WriteConfig()
}
