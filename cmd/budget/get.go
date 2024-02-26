package budget

import (
	"github.com/ibilalkayy/flow/internal/app"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the budget data in CSV",
	Run: func(cmd *cobra.Command, args []string) {
		filepath, _ := cmd.Flags().GetString("filepath")
		filename, _ := cmd.Flags().GetString("filename")
		app.GetBudgetData(filepath, filename)
	},
}

func init() {
	getCmd.Flags().StringP("filepath", "p", "", "Give the file path to store the data")
	getCmd.Flags().StringP("filename", "n", "", "Give the CSV file name to store the data")
}
