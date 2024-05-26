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

// AmountCmd represents the total amount command
var AmountCmd = &cobra.Command{
	Use:   "amount",
	Short: "View the total amount",
	Run: func(cmd *cobra.Command, args []string) {

		myConnection := &db.MyConnection{}
		myTotalAmount := &total_amount_db.MyTotalAmountDB{}
		deps := interfaces.Dependencies{
			Connect:     myConnection,
			TotalAmount: myTotalAmount,
		}

		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myTotalAmount.Handler = handle

		table, err := handle.Deps.TotalAmount.ViewTotalAmount()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(table[0])
	},
}
