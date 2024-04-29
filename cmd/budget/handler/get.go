package budget_handler

import (
	"github.com/ibilalkayy/flow/framework_drivers/db/budget_db"
	"github.com/spf13/cobra"
)

// GetCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the budget data in CSV",
	Run: func(cmd *cobra.Command, args []string) {
		filepath, _ := cmd.Flags().GetString("filepath")
		filename, _ := cmd.Flags().GetString("filename")
		budget_db.GetBudgetData(filepath, filename)
	},
}

func init() {
	GetCmd.Flags().StringP("filepath", "p", "", "Give the file path to store the data")
	GetCmd.Flags().StringP("filename", "n", "", "Give the CSV file name to store the data")
}
