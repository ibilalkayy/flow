package total_amount_handler

import (
	"log"

	"github.com/ibilalkayy/flow/internal/framework_drivers/db/total_amount_db"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the total amount data",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		err := total_amount_db.RemoveTotalAmount(category)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RemoveCmd.Flags().StringP("category", "c", "", "Write the category to remove it's date")
}
