package total_amount_subhandler

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/entities"
	"github.com/spf13/cobra"
)

// InactiveCmd represents the inactive command
var InactiveCmd = &cobra.Command{
	Use:   "inactive",
	Short: "Make the total amount inactive",
	Run: func(cmd *cobra.Command, args []string) {
		h := TakeHandler()
		values, err := h.Deps.TotalAmount.ViewTotalAmount()
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
			h.Deps.TotalAmount.UpdateStatus(&updateStatus)
		}
	},
}
