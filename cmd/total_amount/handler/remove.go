package total_amount_handler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/spf13/cobra"
)

func removeData() error {
	err := total_amount_db.RemoveTotalAmount("TotalAmount")
	if err != nil {
		return err
	}

	err = total_amount_db.RemoveTotalAmount("TotalAmountCategory")
	if err != nil {
		return err
	}
	return err
}

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		err := removeData()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Total amount is successfully removed!")
	},
}
