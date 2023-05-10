package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "IPtracker simple CLI app",
		Long:  `IPtracker simple CLI app.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
