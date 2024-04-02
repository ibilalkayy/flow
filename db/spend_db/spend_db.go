package spend_db

import (
	"errors"

	"github.com/ibilalkayy/flow/db"
	"github.com/ibilalkayy/flow/internal/structs"
)

func CreateSpending(sv *structs.SpendingVariables, amountExceeded, basePath string) error {
	data, err := db.Table(basePath, "003_create_spending_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO Spending(categories, category_amounts, spending_amounts, amount_exceeded) VALUES($1, $2, $3, $4)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	if len(sv.Category) != 0 {
		if amountExceeded == "true" {
			_, err = insert.Exec(sv.Category, sv.CategoryAmount, sv.SpendingAmount, "true")
			if err != nil {
				return err
			}
		} else if amountExceeded == "false" {
			_, err = insert.Exec(sv.Category, sv.CategoryAmount, sv.SpendingAmount, "false")
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("category can't be empty")
	}
	return nil
}
