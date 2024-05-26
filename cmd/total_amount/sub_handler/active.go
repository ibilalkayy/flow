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

// ActiveCmd represents the active command
var ActiveCmd = &cobra.Command{
	Use:   "active",
	Short: "Make the total amount active",
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

		if status == "active" {
			fmt.Println(errors.New("the status is already active"))
		} else {
			updateStatus := entities.TotalAmountVariables{
				Status: "Active",
			}
			handle.Deps.TotalAmount.UpdateStatus(&updateStatus)
		}
	},
}
