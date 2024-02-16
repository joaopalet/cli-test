/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"jp-cli/cmd/tables"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type WorldTimeResponse struct {
	Datetime string `json:"datetime"`
}

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "jp-cli",
		Short: fmt.Sprintf("A test CLI build by JP (version %s)", version),
		Long:  fmt.Sprintf("A test CLI build by JP (version %s)", version),
		RunE: func(cmd *cobra.Command, args []string) error {
			url := "http://worldtimeapi.org/api/ip"

			response, err := http.Get(url)
			if err != nil {
				return err
			}
			defer response.Body.Close()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				return err
			}

			var worldTime WorldTimeResponse
			err = json.Unmarshal(body, &worldTime)
			if err != nil {
				return err
			}

			table := tables.NewTable()
			table.SetHeader("FIELD", "VALUE")
			table.AddRow("Datetime", worldTime.Datetime)
			err = table.Display(cmd)
			if err != nil {
				return fmt.Errorf("render table: %w", err)
			}
			return nil
		},
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return rootCmd
}

func ExecuteRootCmd(version string) {
	rootCmd := NewRootCmd(version)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
