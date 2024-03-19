package handler

import (
	"log"

	app "github.com/ibilalkayy/flow/internal/app/budget"
	"github.com/spf13/cobra"
)

// MsgCmd represents the msg command
var MsgCmd = &cobra.Command{
	Use:   "msg",
	Short: "The message of alert notifications",
	Run: func(cmd *cobra.Command, args []string) {
		err := app.AlertMessage()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
}
