package budget_db

import (
	"errors"

	"github.com/ibilalkayy/flow/db"
	"github.com/ibilalkayy/flow/internal/common/functions"
	"github.com/ibilalkayy/flow/internal/common/structs"
)

func InsertHistory(hv *structs.HistoryVariables, basePath string) error {
	data, err := db.Table(basePath, "001_create_budget_table.sql", 1)
	if err != nil {
		return err
	}

	query := "INSERT INTO History(dates, categories, amounts, transaction_ids, blockchains, addresses) VALUES($1, $2, $3, $4, $5, $6)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	includedCategory, value, err := functions.TotalAmountValues()
	if err != nil {
		return err
	}

	totalAmount, ok := value[0].(int)
	if !ok {
		return errors.New("unable to convert to int")
	}

	if len(hv.Category) != 0 && len(includedCategory) != 0 {
		if hv.Amount != 0 && totalAmount != 0 {
			_, err = insert.Exec(hv.Date, hv.Category, hv.Amount, hv.TransactionID, hv.Blockchain, hv.Address)
			if err != nil {
				return err
			}
		} else {
			return errors.New("enter the amount")
		}
	} else {
		return errors.New("enter the category")
	}
	return nil
}
