package spend_handler

import (
	"fmt"

	spend_subhandler "github.com/ibilalkayy/flow/cmd/spend/sub_handler"
	"github.com/spf13/cobra"
)

// HistoryCmd represents the history command
var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Show the transaction history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("no command is written")
	},
}

func init() {
	HistoryCmd.AddCommand(spend_subhandler.ShowCmd)
	HistoryCmd.AddCommand(spend_subhandler.RemoveCmd)
}
