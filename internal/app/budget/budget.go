package app

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ibilalkayy/flow/db/budget_db"
	"github.com/ibilalkayy/flow/internal/structs"
	"github.com/jedib0t/go-pretty/v6/table"
)

func CreateBudget(bv *structs.BudgetVariables, basePath string) error {
	data, err := budget_db.Table(basePath, "001_create_budget_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO Budget(categories, amounts) VALUES($1, $2)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	if len(bv.Category) != 0 {
		_, err = insert.Exec(bv.Category, bv.Amount)
		if err != nil {
			return err
		}
		fmt.Println("Budget data is successfully inserted!")
	} else {
		return errors.New("category can't be empty")
	}
	return nil
}

func ViewBudget(category string) (string, error) {
	// Create a new instance of BudgetVariables to hold the retrieved data
	bv := new(structs.BudgetVariables)

	// Connect to the database
	db, err := budget_db.Connection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Prepare the table writer
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Category", "Amount"})

	// Initialize total amount
	totalAmount := 0.0

	// Query the database based on the provided category
	var rows *sql.Rows
	if len(category) != 0 {
		query := "SELECT categories, amounts FROM Budget WHERE categories=$1"
		rows, err = db.Query(query, category)
	} else {
		query := "SELECT categories, amounts FROM Budget"
		rows, err = db.Query(query)
	}
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Iterate over the rows and add them to the table writer
	for rows.Next() {
		if err := rows.Scan(&bv.Category, &bv.Amount); err != nil {
			return "", err
		}
		// Check if amount is empty
		if bv.Amount != "" {
			// Convert bv.Amount to float64
			amount, err := strconv.ParseFloat(bv.Amount, 64)
			if err != nil {
				return "", err
			}
			tw.AppendRow([]interface{}{bv.Category, amount})
			totalAmount += amount // Accumulate total amount
		}
	}

	// Add total amount row to the table
	tw.AppendFooter(table.Row{"Total Amount", totalAmount})

	// Render the table
	tableRender := "Budget Data\n" + tw.Render()

	return tableRender, nil
}

func RemoveBudget(category string) error {
	db, err := budget_db.Connection()
	if err != nil {
		return err
	}

	query := "DELETE FROM Budget WHERE categories=$1"
	remove, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer remove.Close()

	if len(category) != 0 {
		_, err = remove.Exec(category)
		if err != nil {
			return err
		}
		fmt.Printf("'%s' category is successfully removed!\n", category)
	} else {
		fmt.Println("First enter the category and then remove it")
	}
	return nil
}

func UpdateBudget(old, new, amount string) error {
	var count int
	var query string
	var params []interface{}

	db, err := budget_db.Connection()
	if err != nil {
		return err
	}

	// Check if the old category exists
	err = db.QueryRow("SELECT COUNT(*) FROM Budget WHERE categories = $1", old).Scan(&count)
	if err != nil {
		return err
	}

	// If the old category does not exist, return an error
	if count == 0 {
		return errors.New("'" + old + "'" + " category does not exist")
	}

	if len(new) != 0 && len(amount) != 0 {
		query = "UPDATE Budget SET categories=$1, amounts=$2 WHERE categories=$3"
		params = []interface{}{new, amount, old}
	} else if len(new) != 0 {
		query = "UPDATE Budget SET categories=$1 WHERE categories=$2"
		params = []interface{}{new, old}
	} else if len(amount) != 0 {
		query = "UPDATE Budget SET amounts=$1 WHERE categories=$2"
		params = []interface{}{amount, old}
	} else {
		fmt.Println("No field provided to adjust")
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		return err
	}

	fmt.Println("Your budget category is successfully updated!")
	return nil
}

func GetBudgetData(filepath, filename string) error {
	bv := new(structs.BudgetVariables)
	db, err := budget_db.Connection()
	if err != nil {
		return err
	}

	query := "SELECT categories, amounts FROM Budget"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	file, err := os.Create(filepath + "/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Category", "Amount"}
	if err := writer.Write(header); err != nil {
		return err
	}
	for rows.Next() {
		if err := rows.Scan(&bv.Category, &bv.Amount); err != nil {
			return err
		}

		data := []string{bv.Category, bv.Amount}
		if err := writer.Write(data); err != nil {
			return err
		}

		if err := rows.Err(); err != nil {
			return err
		}
	}
	fmt.Printf("Successfully created a '%s' file in '%s'\n", filename, filepath)
	return nil
}

func BudgetAmount(category string) (string, error) {
	if len(category) == 0 {
		return "", errors.New("category is not entered")
	} else {
		bv := new(structs.BudgetVariables)

		db, err := budget_db.Connection()
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
}

func Frequency(value string) (string, error) {
	switch value {
	case "hourly", "Hourly":
		return "You selected the hourly", nil
	case "daily", "Daily":
		return "You selected the daily", nil
	case "weekly", "Weekly":
		return "You selected the weekly", nil
	case "monthly", "Monthly":
		return "You selected the monthly", nil
	default:
		return "", errors.New("invalid selection. write [ hourly | daily | weekly | monthly ]")
	}
}

func Method(value string) (string, error) {
	switch value {
	case "email", "Email":
		return "You selected an email", nil
	case "cli", "Cli", "CLI":
		return "You selected the CLI", nil
	default:
		return "invalid selection. write [ email | CLI ]", nil
	}
}

func Alert(av structs.AlertVariables) error {
	if len(av.Category) != 0 && len(av.Frequency) != 0 && len(av.Method) != 0 {
		amount, err := BudgetAmount(av.Category)
		if err != nil {
			return err
		}

		frequencyMsg, err := Frequency(av.Frequency)
		if err != nil {
			return err
		}

		method, err := Method(av.Method)
		if err != nil {
			return err
		}

		fmt.Println(amount)
		fmt.Println(frequencyMsg)
		fmt.Println(method)
	} else {
		return errors.New("enter all the alert values")
	}
	return nil
}
