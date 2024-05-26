package total_amount_subhandler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// CategoriesCmd represents the category command
var CategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "View the categories in the total amount",
	Run: func(cmd *cobra.Command, args []string) {
		myConnection := &db.MyConnection{}
		myTotalCategory := &total_amount_db.MyTotalAmountDB{}
		deps := interfaces.Dependencies{
			Connect:             myConnection,
			TotalAmountCategory: myTotalCategory,
		}

		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myTotalCategory.Handler = handle

		categories, _, err := handle.Deps.TotalAmountCategory.ViewTotalAmountCategories()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(categories)
	},
}
