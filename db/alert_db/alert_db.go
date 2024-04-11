package alert_db

import (
	"database/sql"
	"errors"

	"github.com/ibilalkayy/flow/db"
	"github.com/ibilalkayy/flow/internal/common/structs"
	"github.com/jedib0t/go-pretty/v6/table"
)

func CreateAlert(av *structs.AlertVariables, basePath string) error {
	data, err := db.Table(basePath, "002_create_alert_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO Alert(categories, alert_methods, alert_frequencies, alert_days, alert_weekdays, alert_hours, alert_minutes, alert_seconds) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	if len(av.Category) != 0 && len(av.Method) != 0 && len(av.Frequency) != 0 {
		_, err = insert.Exec(av.Category, av.Method, av.Frequency, av.Days, av.Weekdays, av.Hours, av.Minutes, av.Seconds)
		if err != nil {
			return err
		}
	} else {
		return errors.New("enter all the flags")
	}
	return nil
}

func ViewAlert(category string) ([9]interface{}, error) {
	av := new(structs.AlertVariables)

	db, err := db.Connection()
	if err != nil {
		return [9]interface{}{}, err
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Categories", "Methods", "Frequencies", "Days", "Weekdays", "Hours", "Minutes", "Seconds"})

	var rows *sql.Rows
	if len(category) != 0 {
		query := "SELECT categories, alert_methods, alert_frequencies, alert_days, alert_weekdays, alert_hours, alert_minutes, alert_seconds FROM Alert WHERE categories=$1"
		rows, err = db.Query(query, category)
	} else {
		query := "SELECT categories, alert_methods, alert_frequencies, alert_days, alert_weekdays, alert_hours, alert_minutes, alert_seconds FROM Alert"
		rows, err = db.Query(query)
	}
	if err != nil {
		return [9]interface{}{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&av.Category, &av.Method, &av.Frequency, &av.Days, &av.Weekdays, &av.Hours, &av.Minutes, &av.Seconds); err != nil {
			return [9]interface{}{}, err
		}

		tw.AppendRow([]interface{}{av.Category, av.Method, av.Frequency, av.Days, av.Weekdays, av.Hours, av.Minutes, av.Seconds})
		tw.AppendSeparator()
	}
	if err := rows.Err(); err != nil {
		return [9]interface{}{}, err
	}

	tableRender := "Alert Info\n" + tw.Render()

	values := [9]interface{}{tableRender, av.Category, av.Method, av.Frequency, av.Days, av.Weekdays, av.Hours, av.Minutes, av.Seconds}
	return values, nil
}
