package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
	"log"
	"time"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start work time on the provided task",
	Run:   startProgress,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func startProgress(cmd *cobra.Command, args []string) {
	var currentActiveIssueId = viper.GetString("active_issue.id")
	var issues = viper.GetStringSlice("issues")
	var nextActiveIssueId = args[0]

	if nextActiveIssueId == "" {
		log.Println("Unknown issue id provided")
		return
	}

	if currentActiveIssueId == nextActiveIssueId {
		log.Println("Provided issue id is already set as active")
		return
	}

	if !slices.Contains(issues, nextActiveIssueId) {
		log.Println("Cannot start progress on the task that does not belong to you")
		return
	}

	var zagrebTimezone, _ = time.LoadLocation("Europe/Zagreb")
	var timeInZagreb = time.Now().In(zagrebTimezone)
	var formattedTimeInZagreb = timeInZagreb.Format("2006-01-02T15:04:05.000Z0700")

	viper.Set("active_issue.id", nextActiveIssueId)
	viper.Set("active_issue.started_at", formattedTimeInZagreb)
	viper.Set("active_issue.started_at_unix", time.Now().Unix())
	viper.WriteConfig()
}
