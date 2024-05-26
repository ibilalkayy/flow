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

// CheckCmd represents the check command
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the status of the total amount",
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

		values, err := handle.Deps.TotalAmount.ViewTotalAmount()
		if err != nil {
			log.Fatal(err)
		}

		status, ok := values[4].(string)
		if !ok {
			fmt.Println("unable to convert to string")
		}

		fmt.Printf("This is your total amount status: %s\n", status)
	},
}
