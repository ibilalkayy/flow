package total_amount_subhandler

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/framework_drivers/db/total_amount_db"
	"github.com/spf13/cobra"
)

// InactiveCmd represents the inactive command
var InactiveCmd = &cobra.Command{
	Use:   "inactive",
	Short: "Make the total amount inactive",
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

		if status == "inactive" {
			fmt.Println(errors.New("the status is already inactive"))
		} else {
			updateStatus := entities.TotalAmountVariables{
				Status: "Inactive",
			}
			action.UpdateStatus(&updateStatus)
		}
	},
}
