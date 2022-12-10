package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/DooomiT/rudi-go/api"
	"github.com/asticode/go-asticoqui"
	"github.com/spf13/cobra"
)

func ServeLocal(groupId string) *cobra.Command {
	return &cobra.Command{
		Use:     "serve-local <model> [scorer] [port]",
		Long:    "Run this command in order to start a api server with a local stt instance",
		GroupID: groupId,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := serveApiLocal(args)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

}

func getModel(modelpath, scorerpath string) (*asticoqui.Model, error) {
	if _, err := os.Stat(modelpath); err != nil {
		return nil, err
	}
	m, err := asticoqui.New(modelpath)
	if err != nil {
		return nil, fmt.Errorf("failed initializing model: %v", err)
	}
	if scorerpath != "" {
		if _, err := os.Stat(scorerpath); err != nil {
			return nil, err
		}
		if err := m.EnableExternalScorer(scorerpath); err != nil {
			return nil, fmt.Errorf("failed enabling external scorer: %v", err)
		}
	}
	return m, nil
}

func serveApiLocal(args []string) error {
	port := "3000"
	modelpath := args[0]
	scorerpath := ""
	if len(args) == 2 {
		_, err := strconv.Atoi(args[1])
		if err != nil {
			scorerpath = args[1]
		} else {
			port = args[1]
		}
	}
	if len(args) == 3 {
		port = args[2]
	}
	model, err := getModel(modelpath, scorerpath)
	if err != nil {
		return err
	}
	defer model.Close()
	fmt.Printf("Sample Rate: %d", model.SampleRate())
	api := api.NewLocal("v1", model)

	return api.Run(port)
}
