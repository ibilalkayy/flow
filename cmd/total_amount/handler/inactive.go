package total_amount_handler

import (
	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/ibilalkayy/flow/internal/common/structs"
	"github.com/spf13/cobra"
)

// InactiveCmd represents the inactive command
var InactiveCmd = &cobra.Command{
	Use:   "inactive",
	Short: "Make the total amount inactive",
	Run: func(cmd *cobra.Command, args []string) {
		status := structs.TotalAmountVariables{
			Status: "Inactive",
		}
		total_amount_db.UpdateStatus(&status)
	},
}
