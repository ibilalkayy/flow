package internal_spending

import (
	"time"

	"github.com/ibilalkayy/flow/db/budget_db"
	"github.com/ibilalkayy/flow/internal/common/structs"
)

func StoreHistory(category string, spending_amount int) error {
	currentDate := time.Now().Format("2006-01-02")

	hv := structs.HistoryVariables{
		Date:          currentDate,
		Category:      category,
		Amount:        spending_amount,
		TransactionID: "transaction id",
		Blockchain:    "ethereum",
		Address:       "ethereum address",
	}

	err := budget_db.InsertHistory(&hv, "db/migrations/")
	if err != nil {
		return err
	}
	return nil
}
