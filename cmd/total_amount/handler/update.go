package total_amount_handler

import (
	"fmt"

	"github.com/spf13/cobra"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetString("amount")
		label, _ := cmd.Flags().GetString("label")
		fmt.Println(amount)
		fmt.Println(label)
	},
}

func init() {
	UpdateCmd.Flags().StringP("amount", "a", "", "Write the total amount that you want to update")
	UpdateCmd.Flags().StringP("label", "l", "", "Write the total amount that you want to update")
}
