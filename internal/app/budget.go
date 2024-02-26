package app

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/ibilalkayy/flow/db"
	"github.com/jedib0t/go-pretty/v6/table"
)

type BudgetVariables struct {
	Category string
	Amount   string
}

func CreateBudget(bv *BudgetVariables) error {
	data, err := db.Table("budget", "001_create_budget_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO Budget(categories, amounts) VALUES($1, $2)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}

	defer insert.Close()

	_, err = insert.Exec(bv.Category, bv.Amount)
	if err != nil {
		return err
	}
	fmt.Println("Budget data is successfully inserted!")
	return nil
}

func ViewBudget(category string) (string, error) {
	// Create a new instance of BudgetVariables to hold the retrieved data
	bv := new(BudgetVariables)

	// Connect to the database
	db, err := db.Connection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Prepare the table writer
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Category", "Amount"})

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
		tw.AppendRow([]interface{}{bv.Category, bv.Amount})
	}

	// Render the table
	return "Budget Data\n" + tw.Render(), nil
}

func RemoveBudget(category string) error {
	db, err := db.Connection()
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
		fmt.Printf("%s category is successfully removed!", category)
	} else {
		fmt.Println("First enter the category and then remove it")
	}
	return nil
}

func UpdateBudget(old, new, amount string) error {
	var count int
	var query string
	var params []interface{}

	db, err := db.Connection()
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

	if len(new) != 0 {
		query = "UPDATE Budget SET categories=$1 WHERE categories=$2"
		params = []interface{}{new, old}
	} else if len(amount) != 0 {
		query = "UPDATE Budget SET amounts=$1 WHERE categories=$2"
		params = []interface{}{amount, old}
	} else if len(new) != 0 && len(amount) != 0 {
		query = "UPDATE Budget SET categories=$1, amounts=$2 WHERE categories=$3"
		params = []interface{}{new, amount, old}
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
	bv := new(BudgetVariables)
	db, err := db.Connection()
	if err != nil {
		return err
	}

	// var rows *sql.Rows
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
	return nil
}
