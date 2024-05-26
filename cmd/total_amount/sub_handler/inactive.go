package total_amount_subhandler

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	"github.com/spf13/cobra"
)

// InactiveCmd represents the inactive command
var InactiveCmd = &cobra.Command{
	Use:   "inactive",
	Short: "Make the total amount inactive",
	Run: func(cmd *cobra.Command, args []string) {

		myConnection := &db.MyConnection{}
		myTotalDB := &total_amount_db.MyTotalAmountDB{}
		deps := interfaces.Dependencies{
			Connect:     myConnection,
			TotalAmount: myTotalDB,
		}

		handle := handler.NewHandler(deps)
		myConnection.Handler = handle
		myTotalDB.Handler = handle

		values, err := handle.Deps.TotalAmount.ViewTotalAmount()
		if err != nil {
			log.Fatal(err)
		}

		status, ok := values[4].(string)
		if !ok {
			fmt.Println("unable to convert to string")
		}

		if status == "inactive" {
			fmt.Println(errors.New("the status is already inactive"))
		} else {
			updateStatus := entities.TotalAmountVariables{
				Status: "Inactive",
			}
			handle.Deps.TotalAmount.UpdateStatus(&updateStatus)
		}
	},
}
