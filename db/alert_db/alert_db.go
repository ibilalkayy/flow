package alert_db

import (
	"errors"

	"github.com/ibilalkayy/flow/db"
	internal_budget "github.com/ibilalkayy/flow/internal/app/budget"
	"github.com/ibilalkayy/flow/internal/structs"
)

func CreateAlert(av *structs.AlertVariables, basePath string) error {
	data, err := db.Table(basePath, "002_create_alert_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO Alert(categories, category_amounts, alert_methods, alert_frequencies) VALUES($1, $2, $3, $4)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	var category, categoryAmount string

	if len(av.Category) != 0 && len(av.Method) != 0 && len(av.Frequency) != 0 {
		categoryPresent, err := internal_budget.IsCategoryPresent(av.Category)
		if err != nil {
			return err
		}

		if categoryPresent {
			category = av.Category
			categoryAmount, err = internal_budget.CategoryAmount(category)
			if err != nil {
				return err
			}
		} else {
			return errors.New("enter the right category")
		}
	} else {
		return errors.New("enter the required flags")
	}

	_, err = insert.Exec(category, categoryAmount, av.Method, av.Frequency)
	if err != nil {
		return err
	}
	return nil
}

func ViewAlert(category string) ([2]string, error) {
	ev := new(structs.EmailVariables)

	db, err := db.Connection()
	if err != nil {
		return [2]string{}, err
	}

	query := "SELECT categories, category_amounts FROM Alert WHERE categories=$1"
	rows, err := db.Query(query, category)
	if err != nil {
		return [2]string{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&ev.Category, &ev.CategoryAmount); err != nil {
			return [2]string{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return [2]string{}, err
	}

	values := [2]string{ev.Category, ev.CategoryAmount}
	return values, nil
}
