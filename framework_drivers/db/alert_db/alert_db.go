package alert_db

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ibilalkayy/flow/entities"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (m MyAlertDatabase) CreateAlert(av *entities.AlertVariables) error {
	data, err := m.Table("framework_drivers/db/migrations/002_create_alert_table.sql", 0)
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

func (m MyAlertDatabase) ViewAlert(category string) ([9]interface{}, error) {
	av := new(entities.AlertVariables)

	db, err := m.Connection()
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

func (m MyAlertDatabase) RemoveAlert(category string) error {
	db, err := m.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	data, err := m.ViewAlert(category)
	if err != nil {
		return err
	}

	foundCategory, ok := data[1].(string)
	if !ok {
		return errors.New("unable to convert data to string")
	}

	query := "DELETE FROM Alert"
	var args []interface{}

	if len(category) != 0 {
		if len(foundCategory) != 0 {
			query += " WHERE categories=$1"
			args = append(args, category)
		} else {
			return errors.New("category is not found")
		}
	}

	remove, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer remove.Close()

	_, err = remove.Exec(args...)
	if err != nil {
		return err
	}

	if len(category) != 0 {
		fmt.Printf("Alert values of '%s' category is successfully removed", category)
	} else {
		if len(foundCategory) != 0 {
			fmt.Printf("Alert data is successfully deleted!")
		} else {
			return errors.New("no data is found")
		}
	}

	return nil
}

func (m MyAlertDatabase) UpdateAlert(av *entities.AlertVariables) error {
	db, err := m.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Alert WHERE categories = $1", av.Category).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("'" + av.Category + "'" + " category does not exist")
	}

	var params []interface{}
	query := "UPDATE Alert SET "
	paramCount := 1 // Start with $1 for the first parameter

	if len(av.Category) != 0 {
		query += "categories=$" + strconv.Itoa(paramCount) + ", "
		params = append(params, av.Category)
		paramCount++
	}
	if len(av.Method) != 0 {
		query += "alert_methods=$" + strconv.Itoa(paramCount) + ", "
		params = append(params, av.Method)
		paramCount++
	}
	if len(av.Frequency) != 0 {
		query += "alert_frequencies=$" + strconv.Itoa(paramCount) + ", "
		params = append(params, av.Frequency)
		paramCount++
	}
	if av.Days != 0 {
		query += "alert_days=$" + strconv.Itoa(paramCount) + ", "
		params = append(params, av.Days)
		paramCount++
	}
	if len(av.Weekdays) != 0 {
		query += "alert_weekdays=$" + strconv.Itoa(paramCount) + ", "
		params = append(params, av.Weekdays)
		paramCount++
	}
	if av.Hours != 0 {
		query += "alert_hours=$" + strconv.Itoa(paramCount) + ", "
		params = append(params, av.Hours)
		paramCount++
	}
	if av.Minutes != 0 {
		query += "alert_minutes=$" + strconv.Itoa(paramCount) + ", "
		params = append(params, av.Minutes)
		paramCount++
	}
	if av.Seconds != 0 {
		query += "alert_seconds=$" + strconv.Itoa(paramCount) + ", "
		params = append(params, av.Seconds)
		paramCount++
	}

	// Remove the trailing comma and space
	query = strings.TrimSuffix(query, ", ")

	query += " WHERE categories=$" + strconv.Itoa(paramCount)
	params = append(params, av.Category)

	_, err = db.Exec(query, params...)
	if err != nil {
		return err
	}

	fmt.Println("Alert updated successfully!")
	return nil
}
