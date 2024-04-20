package total_amount_subhandler

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/ibilalkayy/flow/internal/common/structs"
	"github.com/spf13/cobra"
)

// ActiveCmd represents the active command
var ActiveCmd = &cobra.Command{
	Use:   "active",
	Short: "Make the total amount active",
	Run: func(cmd *cobra.Command, args []string) {
		values, err := total_amount_db.ViewTotalAmount()
		if err != nil {
			log.Fatal(err)
		}

		status, ok := values[2].(string)
		if !ok {
			fmt.Println("unable to convert to string")
		}

		if status == "active" {
			fmt.Println(errors.New("the status is already active"))
		} else {
			updateStatus := structs.TotalAmountVariables{
				Status: "Active",
			}
			total_amount_db.UpdateStatus(&updateStatus)
		}
	},
}
