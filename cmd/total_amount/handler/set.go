package total_amount_handler

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SetCmd represents the set command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the total amount",
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetString("amount")
		include_category, _ := cmd.Flags().GetString("include")
		exclude_category, _ := cmd.Flags().GetString("exclude")
		label, _ := cmd.Flags().GetString("label")
		fmt.Println(amount)
		fmt.Println(include_category)
		fmt.Println(exclude_category)
		fmt.Println(label)
	},
}

func init() {
	SetCmd.Flags().StringP("amount", "a", "", "Write the total amount that you want to set")
	SetCmd.Flags().StringP("include", "i", "", "Specify a category to include in the total amount")
	SetCmd.Flags().StringP("exclude", "e", "", "Specify a category to exclude in the total amount")
	SetCmd.Flags().StringP("label", "l", "", "Provide a label for setting up your total amount")
}
