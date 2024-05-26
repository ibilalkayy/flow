package total_amount_handler

import (
	"log"

	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")

		myConnection := &db.MyConnection{}
		myTotalBD := &total_amount_db.MyTotalAmountDB{}
		deps := interfaces.Dependencies{
			Connect:             myConnection,
			TotalAmount:         myTotalBD,
			TotalAmountCategory: myTotalBD,
		}

		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myTotalBD.Handler = handle
		err := handle.Deps.TotalAmount.RemoveTotalAmount(category)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RemoveCmd.Flags().StringP("category", "c", "", "Write the category to remove it's date")
}
