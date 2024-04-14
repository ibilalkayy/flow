package total_amount_handler

import (
	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/ibilalkayy/flow/internal/common/structs"
	"github.com/spf13/cobra"
)

// ActiveCmd represents the active command
var ActiveCmd = &cobra.Command{
	Use:   "active",
	Short: "Make the total amount active",
	Run: func(cmd *cobra.Command, args []string) {
		status := structs.TotalAmountVariables{
			Status: "Active",
		}
		total_amount_db.UpdateStatus(&status)
	},
}

func init() {
}
