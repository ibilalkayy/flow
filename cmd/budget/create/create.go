package create

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/db"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the budget of different categories",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		amount, _ := cmd.Flags().GetString("amount")
		fmt.Println(category)
		fmt.Println(amount)
		err := db.Table("budget", "001_create_budget_table.sql", 0)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	CreateCmd.Flags().StringP("category", "c", "", "Write the category like groceries, utilities")
	CreateCmd.Flags().StringP("amount", "a", "", "Write the total amount for that category")
}
