package usecases_budget

import (
	"errors"

	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/handler"
)

type MyBudget struct {
	*handler.Handler
}

func (h MyBudget) CategoryAmount(category string) (string, int, error) {
	bv := new(entities.BudgetVariables)

	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return "", 0, err
	}

	checkQuery := "SELECT COUNT(*) FROM Budget WHERE categories=$1"
	var count int
	err = db.QueryRow(checkQuery, category).Scan(&count)
	if err != nil {
		return "", 0, nil
	}

	if count == 0 {
		return "", 0, errors.New("first create a budget. go to 'flow budget -h' for help")
	}

	query := "SELECT categories, amounts FROM Budget WHERE categories=$1"
	rows, err := db.Query(query, category)
	if err != nil {
		return "", 0, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&bv.Category, &bv.Amount); err != nil {
			return "", 0, err
		}
	}
	if err := rows.Err(); err != nil {
		return "", 0, err
	}

	return bv.Category, bv.Amount, nil
}
