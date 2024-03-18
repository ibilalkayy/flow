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
		fmt.Println("msg is called")
	},
}

func init() {
}
