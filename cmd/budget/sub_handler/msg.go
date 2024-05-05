package budget_subhandler

import (
	"log"

	usecases_alert "github.com/ibilalkayy/flow/usecases/app/alert"
	"github.com/spf13/cobra"
)

// MsgCmd represents the msg command
var MsgCmd = &cobra.Command{
	Use:   "msg",
	Short: "The CLI message for the alert notifications",
	Run: func(cmd *cobra.Command, args []string) {
		var m usecases_alert.MyAlerts

		category, _ := cmd.Flags().GetString("category")
		err := m.CheckNotification(category)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	MsgCmd.Flags().StringP("category", "c", "", "Write the category to get the notification")
}
