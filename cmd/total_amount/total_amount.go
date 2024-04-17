package total_amount

import (
	"fmt"

	"github.com/ibilalkayy/flow/cmd"
	total_amount_handler "github.com/ibilalkayy/flow/cmd/total_amount/handler"
	"github.com/spf13/cobra"
)

// TotalAmountCmd represents the total-amount command
var TotalAmountCmd = &cobra.Command{
	Use:   "total-amount",
	Short: "Manage your total amount",
	Long: `This command allows you to manage your total amount by setting, view, removing
doing category selection, excluding categories, etc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("no proper command is given. see 'flow total-amount -h'")
	},
}

func init() {
	cmd.RootCmd.AddCommand(TotalAmountCmd)
	// Subcommands
	TotalAmountCmd.AddCommand(total_amount_handler.SetCmd)
	TotalAmountCmd.AddCommand(total_amount_handler.UpdateCmd)
	TotalAmountCmd.AddCommand(total_amount_handler.ViewCmd)
	TotalAmountCmd.AddCommand(total_amount_handler.RemoveCmd)
	TotalAmountCmd.AddCommand(total_amount_handler.StatusCmd)
}
