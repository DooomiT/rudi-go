package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rudi",
	Short: "Rudi is a basic api for speech transcription",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	allGroupId := "general"
	userGroupId := "user"
	allGroup := &cobra.Group{ID: allGroupId, Title: allGroupId}
	userGroup := &cobra.Group{ID: userGroupId, Title: userGroupId}
	rootCmd.AddGroup(allGroup, userGroup)
	rootCmd.AddCommand(Serve(allGroupId))
	rootCmd.AddCommand(ServeLocal(allGroupId))
	rootCmd.AddCommand(SpeechToText(userGroupId))
}
