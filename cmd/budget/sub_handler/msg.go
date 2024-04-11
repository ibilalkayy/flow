package budget_subhandler

import (
	internal_alert "github.com/ibilalkayy/flow/internal/app/alert"
	"github.com/spf13/cobra"
)

// MsgCmd represents the msg command
var MsgCmd = &cobra.Command{
	Use:   "msg",
	Short: "The CLI message for the alert notifications",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		internal_alert.CheckNotification(category)
	},
}

func init() {
	MsgCmd.Flags().StringP("category", "c", "", "Write the category to get the notification")
}
