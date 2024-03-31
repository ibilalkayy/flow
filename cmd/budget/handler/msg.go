package handler

import (
	"fmt"
	"log"

	internal_spending "github.com/ibilalkayy/flow/internal/app/spend"
	"github.com/spf13/cobra"
)

// MsgCmd represents the msg command
var MsgCmd = &cobra.Command{
	Use:   "msg",
	Short: "The CLI message for alert notifications",
	Run: func(cmd *cobra.Command, args []string) {
		err := internal_spending.AlertFrequency("second")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("msg is printed")
	},
}

func init() {
}
