package total_amount_subhandler

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// CategoriesCmd represents the category command
var CategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "View the categories in the total amount",
	Run: func(cmd *cobra.Command, args []string) {
		h := TakeHandler()
		categories, _, err := h.Deps.TotalAmountCategory.ViewTotalAmountCategories()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(categories)
	},
}
