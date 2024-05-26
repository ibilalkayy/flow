package budget_db

import "github.com/ibilalkayy/flow/entities"

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
