package budget

import (
	"fmt"

	"github.com/ibilalkayy/flow/cmd"
	"github.com/ibilalkayy/flow/cmd/budget/handler"
	"github.com/ibilalkayy/flow/email"
	"github.com/spf13/cobra"
)

// budgetCmd represents the budget command
var budgetCmd = &cobra.Command{
	Use:   "budget",
	Short: "Manage your budget",
	Long: `Budget command allows users to manage their budgetary allocations 
for different spending categories. With this command, you can create, view, 
and adjust their budgets to effectively track and control their expenses.`,

	Run: func(cmd *cobra.Command, args []string) {
		email.SendAlertMail()
		fmt.Println("Nothing specified, nothing added")
	},
}

func init() {
	cmd.RootCmd.AddCommand(budgetCmd)
	// Added subcommands
	budgetCmd.AddCommand(handler.CreateCmd)
	budgetCmd.AddCommand(handler.ViewCmd)
	budgetCmd.AddCommand(handler.AdjustCmd)
	budgetCmd.AddCommand(handler.RemoveCmd)
	budgetCmd.AddCommand(handler.GetCmd)
	budgetCmd.AddCommand(handler.AlertCmd)
}
