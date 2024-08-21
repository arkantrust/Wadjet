package cmd

import (
	"fmt"
	"time"

	"github.com/arkantrust/wadjet/cli/transport"
	"github.com/spf13/cobra"
)

// time: Returns the server's local time.
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "The 'time' subcommand will output the specified server's local time.",
	Long: `The 'time' subcommand requests the specified server's local time. For example:

'wadjet time --host localhost --port 6000'.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		Time(host, port)
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.Flags().StringP("host", "H", "localhost", "The server's hostname or IP address.")
	timeCmd.Flags().StringP("port", "p", "6000", "The server's port.")
}

// Time returns the server's local time.
func Time(host string, port string) {
	res, err := transport.Send(host, port, "time")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// parse the time
	t, err := time.Parse(time.RFC3339, res)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(t)
}
