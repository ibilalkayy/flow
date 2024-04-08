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
		active, _ := cmd.Flags().GetString("active")
		inactive, _ := cmd.Flags().GetString("inactive")
		fmt.Println(active)
		fmt.Println(inactive)
	},
}

func init() {
	cmd.RootCmd.AddCommand(TotalAmountCmd)
	// Subcommands
	TotalAmountCmd.AddCommand(total_amount_handler.SetCmd)
	TotalAmountCmd.AddCommand(total_amount_handler.UpdateCmd)
	TotalAmountCmd.AddCommand(total_amount_handler.ViewCmd)
	TotalAmountCmd.AddCommand(total_amount_handler.RemoveCmd)
	// Flags
	TotalAmountCmd.Flags().StringP("active", "a", "", "Make the total amount active")
	TotalAmountCmd.Flags().StringP("inactive", "i", "", "Make the total amount inactive")
}
