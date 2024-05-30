package budget_db

import (
	"errors"

	"github.com/ibilalkayy/flow/entities"
)

func (h MyBudgetDB) TakeBudgetAmount() ([]string, []int, error) {
	bv := new(entities.BudgetVariables)
	var amounts []int
	var categories []string

	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	query := "SELECT categories, amounts FROM Budget"
	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&bv.Category, &bv.Amount)
		if err != nil {
			return nil, nil, err
		}
		amounts = append(amounts, bv.Amount)
		categories = append(categories, bv.Category)
	}

	return categories, amounts, nil
}

func (h MyBudgetDB) BudgetAmountWithException(bv *entities.BudgetVariables) (int, error) {
	var amounts int

	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := "SELECT amounts FROM Budget WHERE NOT (categories=$1)"
	rows, err := db.Query(query, &bv.Category)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&amounts)
		if err != nil {
			return 0, err
		}
	}
	return amounts, nil
}

func (MyBudgetDB) CalculateRemaining(cr *entities.BudgetCalculateVariables) ([2]int, error) {
	// If the new budget amount is greater than or less than the amount in the database
	if cr.BudgetAmount != cr.BudgetAmountInDB {
		if cr.SpentAmountInDB > cr.BudgetAmount {
			// If spent amount is greater than the new budget, reset both spent and remaining amounts to 0
			cr.SpentAmountInDB = 0
			cr.RemainingAmountInDB = 0
		} else {
			// Calculate the remaining amount based on the new budget and spent amount
			cr.RemainingAmountInDB = cr.BudgetAmount - cr.SpentAmountInDB
		}
	} else {
		// If the new budget amount is the same as the old one, return an error
		return [2]int{}, errors.New("this amount is already present. enter a different amount")
	}

	result := [2]int{cr.SpentAmountInDB, cr.RemainingAmountInDB}
	return result, nil
}
