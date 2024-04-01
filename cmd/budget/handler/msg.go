package handler

import (
	"fmt"

	"github.com/spf13/cobra"
)

// MsgCmd represents the msg command
var MsgCmd = &cobra.Command{
	Use:   "msg",
	Short: "The CLI message for alert notifications",
	Run: func(cmd *cobra.Command, args []string) {
		// err := internal_spending.AlertFrequency("second")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		fmt.Println("msg is printed")
	},
}

func init() {
}
