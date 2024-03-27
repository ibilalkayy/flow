package handler

import (
	"fmt"

	"github.com/spf13/cobra"
)

// MsgCmd represents the msg command
var MsgCmd = &cobra.Command{
	Use:   "msg",
	Short: "The message of alert notifications",
	Run: func(cmd *cobra.Command, args []string) {
		// err := internal_alert.AlertMessage()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		fmt.Println("msg is printed")
	},
}

func init() {
}
