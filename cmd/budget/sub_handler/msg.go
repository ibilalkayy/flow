package budget_subhandler

import (
	"log"

	"github.com/spf13/cobra"
)

// MsgCmd represents the msg command
var MsgCmd = &cobra.Command{
	Use:   "msg",
	Short: "The CLI message for the alert notifications",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")

		h := TakeHandler()
		err := h.Deps.ManageAlerts.SendNotification(category)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	MsgCmd.Flags().StringP("category", "c", "", "Write the category to get the notification")
}
