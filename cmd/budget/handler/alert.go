package budget_handler

import (
	"fmt"

	budget_subhandler "github.com/ibilalkayy/flow/cmd/budget/sub_handler"
	"github.com/spf13/cobra"
)

// AlertCmd represents the alert command
var AlertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Get notification once you pass the budget",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Nothing specified, nothing added")
	},
}

func init() {
	AlertCmd.AddCommand(budget_subhandler.MsgCmd)
	AlertCmd.AddCommand(budget_subhandler.SetupCmd)
	AlertCmd.AddCommand(budget_subhandler.ViewCmd)
	AlertCmd.AddCommand(budget_subhandler.UpdateCmd)
	AlertCmd.AddCommand(budget_subhandler.RemoveCmd)
}
