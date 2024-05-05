package total_amount_subhandler

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/framework_drivers/db/total_amount_db"
	"github.com/spf13/cobra"
)

// ActiveCmd represents the active command
var ActiveCmd = &cobra.Command{
	Use:   "active",
	Short: "Make the total amount active",
	Run: func(cmd *cobra.Command, args []string) {
		var action total_amount_db.MyTotalDatabase
		values, err := action.ViewTotalAmount()
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
			action.UpdateStatus(&updateStatus)
		}
	},
}
