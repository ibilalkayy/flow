package total_amount_handler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/spf13/cobra"
)

// CategoriesCmd represents the category command
var CategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "View the categories included in the total amount",
	Run: func(cmd *cobra.Command, args []string) {
		categories, _, err := total_amount_db.ViewTotalAmountCategory()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(categories)
	},
}
