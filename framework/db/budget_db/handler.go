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

func (h MyBudgetDB) ListOfExpection(bv *entities.BudgetVariables) ([]int, error) {
	var spents []int

	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT spent FROM Budget WHERE NOT (categories=$1)"
	rows, err := db.Query(query, &bv.Category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var spent int
		err := rows.Scan(&spent)
		if err != nil {
			return nil, err
		}
		spents = append(spents, spent)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return spents, nil
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

func (h MyBudgetDB) CalculateRemaining(cr *entities.BudgetCalculateVariables) ([2]int, error) {
	// If the new budget amount is greater than or less than the amount in the database
	if cr.BudgetAmount != cr.BudgetAmountInDB {
		// If spent amount is greater than the new budget, reset both spent and remaining amounts to 0
		if cr.SpentAmountInDB > cr.BudgetAmount {
			cr.SpentAmountInDB = 0
			cr.RemainingAmountInDB = cr.BudgetAmount

			values, err := h.Deps.TotalAmount.ViewTotalAmount()
			if err != nil {
				return [2]int{}, err
			}

			totalAmount, ok := values[1].(int)
			if !ok {
				return [2]int{}, errors.New("unable to convert to int")
			}

			category := entities.BudgetVariables{Category: cr.Category}
			spentAmounts, err := h.Deps.ManageBudget.ListOfExpection(&category)
			if err != nil {
				return [2]int{}, err
			}

			totalSpent := 0
			for _, spentAmount := range spentAmounts {
				totalSpent += spentAmount
			}

			if totalSpent != 0 {
				err = h.Deps.TotalAmount.UpdateSpentAndRemaining(totalSpent, totalAmount-totalSpent)
				if err != nil {
					return [2]int{}, err
				}
			} else {
				err = h.Deps.TotalAmount.UpdateSpentAndRemaining(0, 0)
				if err != nil {
					return [2]int{}, err
				}
			}
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
