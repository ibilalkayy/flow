package spend_subhandler

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/framework_drivers/db/budget_db"
	"github.com/spf13/cobra"
)

// ShowCmd represents the show command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the history data",
	Run: func(cmd *cobra.Command, args []string) {
		var m budget_db.MyHistoryDatabase

		category, _ := cmd.Flags().GetString("category")
		table, err := m.ViewHistory(category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(table[0])
	},
}

func init() {
	ShowCmd.Flags().StringP("category", "c", "", "Write the category to show it's history")
}
