package app

import (
	"fmt"

	"github.com/ibilalkayy/flow/db"
)

type BudgetVariables struct {
	Category string
	Amount   string
}

func CreateBudget(bv *BudgetVariables) error {
	data, err := db.Table("budget", "001_create_budget_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO Budget(categories, amounts) VALUES($1, $2)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}

	defer insert.Close()

	_, err = insert.Exec(bv.Category, bv.Amount)
	if err != nil {
		return err
	}
	fmt.Println("Budget data is successfully inserted!")
	return nil
}
