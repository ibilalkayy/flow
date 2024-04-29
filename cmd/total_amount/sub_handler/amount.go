package total_amount_subhandler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/framework_drivers/db/total_amount_db"
	"github.com/spf13/cobra"
)

// AmountCmd represents the total amount command
var AmountCmd = &cobra.Command{
	Use:   "amount",
	Short: "View the total amount",
	Run: func(cmd *cobra.Command, args []string) {
		table, err := total_amount_db.ViewTotalAmount()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(table[0])
	},
}
