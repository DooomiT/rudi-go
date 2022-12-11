package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DooomiT/rudi-go/pkg/rudi"

	"github.com/spf13/cobra"
)

func SpeechToText(groupId string) *cobra.Command {
	return &cobra.Command{
		Use:     "stt [filepath]",
		Long:    "This command uploads a audio file and returns the transcribed text",
		GroupID: groupId,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := stt(args[0])
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

}

func stt(filepath string) error {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	client := rudi.New(http.Client{}, "http://localhost:3000")
	result, err := client.SpeechToText(content)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
