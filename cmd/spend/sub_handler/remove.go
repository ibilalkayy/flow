package spend_subhandler

import (
	"log"

	"github.com/ibilalkayy/flow/framework_drivers/db/budget_db"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the history data",
	Run: func(cmd *cobra.Command, args []string) {
		var m budget_db.MyHistoryDatabase

		category, _ := cmd.Flags().GetString("category")
		err := m.RemoveHistory(category)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RemoveCmd.Flags().StringP("category", "c", "", "Write the category to remove it's history")
}
