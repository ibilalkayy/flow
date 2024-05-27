package total_amount_subhandler

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/entities"
	"github.com/spf13/cobra"
)

// ActiveCmd represents the active command
var ActiveCmd = &cobra.Command{
	Use:   "active",
	Short: "Make the total amount active",
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

		if status == "active" {
			fmt.Println(errors.New("the status is already active"))
		} else {
			updateStatus := entities.TotalAmountVariables{
				Status: "Active",
			}
			h.Deps.TotalAmount.UpdateStatus(&updateStatus)
		}
	},
}
