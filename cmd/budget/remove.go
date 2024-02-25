package budget

import (
	"log"

	"github.com/ibilalkayy/flow/internal/app"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the budget details",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		err := app.RemoveBudget(category)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	removeCmd.Flags().StringP("category", "c", "", "Write the category name to remove")
}
