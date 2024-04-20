package total_amount_subhandler

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/ibilalkayy/flow/internal/common/structs"
	"github.com/spf13/cobra"
)

// InactiveCmd represents the inactive command
var InactiveCmd = &cobra.Command{
	Use:   "inactive",
	Short: "Make the total amount inactive",
	Run: func(cmd *cobra.Command, args []string) {
		values, err := total_amount_db.ViewTotalAmount()
		if err != nil {
			log.Fatal(err)
		}

		status, ok := values[2].(string)
		if !ok {
			fmt.Println("unable to convert to string")
		}

		if status == "inactive" {
			fmt.Println(errors.New("the status is already inactive"))
		} else {
			updateStatus := structs.TotalAmountVariables{
				Status: "Inactive",
			}
			total_amount_db.UpdateStatus(&updateStatus)
		}
	},
}
