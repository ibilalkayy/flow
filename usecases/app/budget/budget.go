package usecases_budget

import (
	"errors"

	"github.com/ibilalkayy/flow/entities"
)

func (m MyBudget) CategoryAmount(category string) (string, int, error) {
	bv := new(entities.BudgetVariables)

	db, err := m.Connection()
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
		return "", 0, errors.New("category not found in the budget")
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
