package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version",
	Long:  `Get the version.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				pip := net.ParseIP(ip)

				if pip == nil {
					log.Fatalln("Invalid IP address")
				}

				showData(pip)
			}
		} else {
			fmt.Println("v0.0.0")
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
