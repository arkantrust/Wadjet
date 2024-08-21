package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wadjet",
	Short: "wadjet-cli is a NetOps server utility.",
	Long: `A Fast and Flexible NetOps server utility built with
				  love by arkantrust in Go.
				  Complete documentation is available at https://wadjet.ddulce.app`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
