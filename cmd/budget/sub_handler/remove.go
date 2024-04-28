package budget_subhandler

import (
	"log"

	"github.com/ibilalkayy/flow/internal/framework_drivers/db/alert_db"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the alert values",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		err := alert_db.RemoveAlert(category)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RemoveCmd.Flags().StringP("category", "c", "", "Write the category name to remove its alert notification values")
}
