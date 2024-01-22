/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd(version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "jp-cli",
		Short: fmt.Sprintf("A test CLI build by JP (version %s)", version),
		Long:  fmt.Sprintf("A test CLI build by JP (version %s)", version),
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
