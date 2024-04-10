package total_amount_handler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/spf13/cobra"
)

// ViewCmd represents the view command
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		table, err := total_amount_db.ViewTotalAmount()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(table[0])
	},
}

func init() {
}
