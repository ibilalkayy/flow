package usecases_spending

import (
	"time"

	"github.com/ibilalkayy/flow/entities"
)

func (h MySpending) StoreHistory(category, recipientAddress string, spending_amount int) error {
	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("03:04:05 PM")

	hv := entities.HistoryVariables{
		Date:             currentDate,
		Time:             currentTime,
		Category:         category,
		Amount:           spending_amount,
		Blockchain:       "ethereum",
		RecipientAddress: recipientAddress,
	}

	err := h.Deps.ManageBudget.InsertHistory(&hv)
	if err != nil {
		return err
	}
	return nil
}
