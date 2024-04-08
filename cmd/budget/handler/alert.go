package budget_handler

import (
	"fmt"

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
	AlertCmd.AddCommand(MsgCmd)
	AlertCmd.AddCommand(SetupCmd)
}
