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

	query := "INSERT INTO Alert(categories, category_amounts, alert_methods, alert_frequencies, alert_days, alert_weekdays, alert_hours, alert_minutes, alert_seconds) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	var category, categoryAmount string

	if len(av.Category) != 0 && len(av.Method) != 0 && len(av.Frequency) != 0 {
		category = av.Category
		categoryAmount, err = internal_budget.CategoryAmount(category)
		if err != nil {
			return err
		}

		_, err = insert.Exec(category, categoryAmount, av.Method, av.Frequency, av.Days, av.Weekdays, av.Hours, av.Minutes, av.Seconds)
		if err != nil {
			return err
		}
	} else {
		return errors.New("enter all the flags")
	}
	return nil
}

func ViewAlert(category string) ([9]string, error) {
	av := new(structs.AlertVariables)

	db, err := db.Connection()
	if err != nil {
		return [9]string{}, err
	}

	query := "SELECT categories, category_amounts, alert_methods, alert_frequencies, alert_days, alert_weekdays, alert_hours, alert_minutes, alert_seconds FROM Alert WHERE categories=$1"
	rows, err := db.Query(query, category)
	if err != nil {
		return [9]string{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&av.Category, &av.CategoryAmount, &av.Method, &av.Frequency, &av.Days, &av.Weekdays, &av.Hours, &av.Minutes, &av.Seconds); err != nil {
			return [9]string{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return [9]string{}, err
	}

	values := [9]string{av.Category, av.CategoryAmount, av.Method, av.Frequency, av.Days, av.Weekdays, av.Hours, av.Minutes, av.Seconds}
	return values, nil
}
