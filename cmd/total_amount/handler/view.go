package total_amount_handler

import (
	"fmt"

	total_amount_subhandler "github.com/ibilalkayy/flow/cmd/total_amount/sub_handler"
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
	ViewCmd.AddCommand(total_amount_subhandler.AmountCmd)
	ViewCmd.AddCommand(total_amount_subhandler.CategoriesCmd)
}
