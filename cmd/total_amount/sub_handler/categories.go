package total_amount_subhandler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/framework_drivers/db/total_amount_db"
	"github.com/spf13/cobra"
)

// CategoriesCmd represents the category command
var CategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "View the categories in the total amount",
	Run: func(cmd *cobra.Command, args []string) {
		var m total_amount_db.MyTotalDatabase

		categories, _, err := m.ViewTotalAmountCategory()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(categories)
	},
}
