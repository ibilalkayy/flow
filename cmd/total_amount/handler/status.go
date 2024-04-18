package total_amount_handler

import (
	"fmt"

	total_amount_subhandler "github.com/ibilalkayy/flow/cmd/total_amount/sub_handler"
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
	StatusCmd.AddCommand(total_amount_subhandler.ActiveCmd)
	StatusCmd.AddCommand(total_amount_subhandler.InactiveCmd)
}
