package internal_budget

import (
	"errors"

	"github.com/ibilalkayy/flow/db"
	"github.com/ibilalkayy/flow/internal/structs"
)

func CategoryAmount(category string) (string, error) {
	bv := new(structs.BudgetVariables)

	db, err := db.Connection()
	if err != nil {
		return "", err
	}

	checkQuery := "SELECT COUNT(*) FROM Budget WHERE categories=$1"
	var count int
	err = db.QueryRow(checkQuery, category).Scan(&count)
	if err != nil {
		return "", nil
	}

	if count == 0 {
		return "", errors.New("category not found")
	}

	query := "SELECT amounts FROM Budget WHERE categories=$1"
	rows, err := db.Query(query, category)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&bv.Amount); err != nil {
			return "", err
		}
	}
	if err := rows.Err(); err != nil {
		return "", err
	}

	return bv.Amount, nil
}

func IsCategoryPresent(category string) (bool, error) {
	db, err := db.Connection()
	if err != nil {
		return false, err
	}

	var count int
	query := "SELECT COUNT(*) FROM Budget WHERE categories=$1"
	err = db.QueryRow(query, category).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
