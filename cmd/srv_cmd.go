package cmd

import (
	"fmt"
	"os"
	"strconv"

	assemblyai "github.com/DooomiT/assembly-ai-go/pkg"
	"github.com/DooomiT/rudi-go/api"
	"github.com/spf13/cobra"
)

func Serve(groupId string) *cobra.Command {
	return &cobra.Command{
		Use:     "serve <assembly-ai-token> [port]",
		Long:    "Run this command in order to start a api server",
		GroupID: groupId,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := serveApi(args)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

}

func maxRetriesFromEnv(def uint) uint {
	maxRetriesEnv := os.Getenv("MAX_RETRIES")
	if maxRetriesEnv != "" {
		mti, err := strconv.Atoi(maxRetriesEnv)
		if err != nil {
			fmt.Println("Environment MAX_RETRIES has to be a number, using default", def)
			return def
		}
		return uint(mti)
	}
	return def
}

func serveApi(args []string) error {
	port := "3000"
	token := args[0]
	client := assemblyai.New("https://api.assemblyai.com/v2", token, nil)
	api := api.New("v1", client, maxRetriesFromEnv(50))
	if len(args) == 2 {
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("[port] has to be a number, you provided: %s", args[0])
		}
		port = args[0]
	}
	return api.Run(port)
}
