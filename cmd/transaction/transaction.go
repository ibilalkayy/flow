package handler

import (
	"fmt"

	"github.com/ibilalkayy/flow/cmd"
	"github.com/spf13/cobra"
)

// transactionCmd represents the transaction command
var TransactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Transaction service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("transaction is called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(TransactionCmd)
}
