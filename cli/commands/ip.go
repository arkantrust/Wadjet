// ip: Returns the server's public IP address.
package cmd

import (
	"fmt"

	"github.com/arkantrust/wadjet/cli/transport"
	"github.com/spf13/cobra"
)

// ip: Returns the server's public IP address.
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "The 'ip' subcommand will output the specified server's public IP address.",
	Long: `The 'ip' subcommand requests the specified server's public IP address. For example:

'wadjet ip --host localhost --port 6000'.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		IP(host, port)
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
	ipCmd.Flags().StringP("host", "H", "localhost", "The server's hostname or IP address.")
	ipCmd.Flags().StringP("port", "p", "6000", "The server's port.")
}

// IP returns the server's public IP address.
func IP(host string, port string) {
	res, err := transport.Send(host, port, "ip")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(res)
}
