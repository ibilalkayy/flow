package budget

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/internal/app"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		details, err := app.ViewBudget(category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Budget Details:\n\tCategory: %s\n\tAmount: %s", details[0], details[1])
	},
}

func init() {
	viewCmd.Flags().StringP("category", "c", "", "Write the category name to show the specific details")
}
