package total_amount_handler

import (
	"log"

	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		err := total_amount_db.RemoveTotalAmount()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
}
