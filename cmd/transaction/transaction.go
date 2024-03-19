package transaction

import (
	"fmt"

	"github.com/ibilalkayy/flow/cmd"
	"github.com/spf13/cobra"
)

func TransactionAmount() int {
	amount := 1600
	return amount
}

// transactionCmd represents the transaction command
var TransactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Transaction service",
	Run: func(cmd *cobra.Command, args []string) {
		amount := TransactionAmount()
		fmt.Println(amount)
	},
}

func init() {
	cmd.RootCmd.AddCommand(TransactionCmd)
}
