package budget_db

import (
	"errors"

	"github.com/ibilalkayy/flow/entities"
)

func (h MyBudgetDB) TakeBudgetAmount() ([]int, error) {
	bv := new(entities.BudgetVariables)
	var amounts []int

	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT amounts FROM Budget"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&bv.Amount)
		if err != nil {
			return nil, err
		}
		amounts = append(amounts, bv.Amount)
	}

	return amounts, nil
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

func (MyBudgetDB) CalculateRemaining(details [4]int) ([2]int, error) {
	if details[0] > details[1] {
		updatedRemaining := details[0] - details[1]
		details[3] += updatedRemaining
	} else if details[0] < details[1] {
		if details[2] <= details[0] {
			details[3] = details[0] - details[2]
		} else {
			details[2] = details[0]
			details[3] = 0
		}
	} else {
		return [2]int{}, errors.New("this amount is already present. enter a different amount")
	}
	result := [2]int{details[2], details[3]}
	return result, nil
}
