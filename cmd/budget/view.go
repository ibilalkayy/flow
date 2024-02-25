package budget

import (
	"fmt"

	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		fmt.Println(category)
	},
}

func init() {
	viewCmd.Flags().StringP("category", "c", "", "Write the category name to show the specific details")
}
