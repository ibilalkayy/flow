package spend_subhandler

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// ShowCmd represents the show command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the history data",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")

		h := TakeHandler()
		table, err := h.Deps.ManageBudget.ViewHistory(category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(table[0])
	},
}

func init() {
	ShowCmd.Flags().StringP("category", "c", "", "Write the category to show it's history")
}
