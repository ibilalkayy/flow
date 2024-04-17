package total_amount_handler

import (
	"fmt"

	"github.com/spf13/cobra"
)

// StatusCmd represents the status command
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Handle the total amount status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("no command is entered. see 'flow total-amount status -h'")
	},
}

func init() {
	StatusCmd.AddCommand(ActiveCmd)
	StatusCmd.AddCommand(InactiveCmd)
}
