package total_amount_handler

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ViewCmd represents the view command
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View your total amount details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("no command is selected. see 'flow total-amount view -h'")
	},
}

func init() {
	ViewCmd.AddCommand(AmountCmd)
	ViewCmd.AddCommand(CategoriesCmd)
}
