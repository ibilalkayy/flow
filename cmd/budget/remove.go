package budget

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		fmt.Println(category)
	},
}

func init() {
	removeCmd.Flags().StringP("category", "c", "", "Write the category name to remove")
}
