package budget_handler

import (
	internal_alert "github.com/ibilalkayy/flow/internal/app/alert"
	"github.com/spf13/cobra"
)

// MsgCmd represents the msg command
var MsgCmd = &cobra.Command{
	Use:   "msg",
	Short: "The CLI message for alert notifications",
	Run: func(cmd *cobra.Command, args []string) {
		internal_alert.CheckNotification("first")
	},
}

func init() {
}
