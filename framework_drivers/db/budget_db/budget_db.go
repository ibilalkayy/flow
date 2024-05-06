package budget_db

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/ibilalkayy/flow/entities"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (m MyBudgetDatabase) CreateBudget(bv *entities.BudgetVariables) error {
	data, err := m.Table("framework_drivers/db/migrations/001_create_budget_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO Budget(categories, amounts, spent, remaining) VALUES($1, $2, $3, $4)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	includedCategory, value, err := m.TotalAmountValues()
	if err != nil {
		return err
	}

	totalAmount, ok := value[0].(int)
	if !ok {
		return errors.New("unable to convert to int")
	}

	for i := 0; i < len(includedCategory); i++ {
		if len(bv.Category) != 0 && len(includedCategory) != 0 && includedCategory[i][0] == bv.Category && totalAmount != 0 {
			_, err = insert.Exec(bv.Category, bv.Amount, 0, 0)
			if err != nil {
				return err
			}
			fmt.Println("Budget data is successfully inserted!")
			return nil
		}
	}
	return errors.New("enter the category in the total amount. see 'flow total-amount -h'")
}

func (m MyBudgetDatabase) ViewBudget(category string) ([5]interface{}, error) {
	// Create a new instance of BudgetVariables to hold the retrieved data
	bv := new(entities.BudgetVariables)

	// Connect to the database
	db, err := m.Connection()
	if err != nil {
		return [5]interface{}{}, err
	}
	defer db.Close()

	// Prepare the table writer
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Category", "Amount", "Spent", "Remaining"})

	// Initialize total amount
	totalAmount := 0

	// Query the database based on the provided category
	var rows *sql.Rows
	if len(category) != 0 {
		query := "SELECT categories, amounts, spent, remaining FROM Budget WHERE categories=$1"
		rows, err = db.Query(query, category)
	} else {
		query := "SELECT categories, amounts, spent, remaining FROM Budget"
		rows, err = db.Query(query)
	}
	if err != nil {
		return [5]interface{}{}, err
	}
	defer rows.Close()

	// Iterate over the rows and add them to the table writer
	for rows.Next() {
		if err := rows.Scan(&bv.Category, &bv.Amount, &bv.Spent, &bv.Remaining); err != nil {
			return [5]interface{}{}, err
		}
		// Check if amount is empty
		if bv.Amount != 0 {
			tw.AppendRow([]interface{}{bv.Category, bv.Amount, bv.Spent, bv.Remaining})
			tw.AppendSeparator()
			totalAmount += bv.Amount
		}
	}

	// Add total amount row to the table
	tw.AppendFooter(table.Row{"Total Amount", totalAmount})

	// Render the table
	tableRender := "Budget Data\n" + tw.Render()

	details := [5]interface{}{tableRender, bv.Category, bv.Amount, bv.Spent, bv.Remaining}
	return details, nil
}

func (m MyBudgetDatabase) RemoveBudget(category string) error {
	db, err := m.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	data, err := m.ViewBudget(category)
	if err != nil {
		return err
	}

	foundCategory, ok := data[1].(string)
	if !ok {
		return errors.New("unable to convert data to string")
	}

	query := "DELETE FROM Budget"
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
		fmt.Printf("'%s' category is successfully removed!\n", category)
	} else {
		if len(foundCategory) != 0 {
			fmt.Printf("Budget data is successfully deleted!")
		} else {
			return errors.New("no data is found")
		}
	}

	return nil
}

func (m MyBudgetDatabase) UpdateBudget(old, new string, amount int) error {
	var count int
	var query string
	var params []interface{}

	db, err := m.Connection()
	if err != nil {
		return err
	}

	// Check if the old category exists
	err = db.QueryRow("SELECT COUNT(*) FROM Budget WHERE categories = $1", old).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("'" + old + "'" + " category does not exist")
	}

	includedCategory, value, err := m.TotalAmountValues()
	if err != nil {
		return err
	}

	totalAmount, ok := value[0].(int)
	if !ok {
		return errors.New("unable to convert to int")
	}

	if len(includedCategory) != 0 && totalAmount != 0 {
		if len(new) != 0 && amount != 0 {
			query = "UPDATE Budget SET categories=$1, amounts=$2 WHERE categories=$3"
			params = []interface{}{new, amount, old}
		} else if len(new) != 0 {
			query = "UPDATE Budget SET categories=$1 WHERE categories=$2"
			params = []interface{}{new, old}
		} else if amount != 0 {
			query = "UPDATE Budget SET amounts=$1 WHERE categories=$2"
			params = []interface{}{amount, old}
		} else {
			fmt.Println("No field provided to adjust")
		}

		_, err = db.Exec(query, params...)
		if err != nil {
			return err
		}
	} else {
		return errors.New("first enter the total amount. see 'flow total-amount -h'")
	}
	return nil
}

func (m MyBudgetDatabase) AddBudgetExpenditure(spent int, category string) error {
	db, err := m.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	var savedSpent, savedRemaining, totalAmount int
	if len(category) != 0 {
		query := "SELECT amounts, spent, remaining FROM Budget WHERE categories = $1"
		err := db.QueryRow(query, category).Scan(&totalAmount, &savedSpent, &savedRemaining)
		if err != nil {
			return err
		}
	} else {
		return errors.New("category is not present")
	}

	totalSpent := spent + savedSpent
	remainingBalance := totalAmount - totalSpent

	includedCategory, value, err := m.TotalAmountValues()
	if err != nil {
		return err
	}

	totalAmount, ok := value[0].(int)
	if !ok {
		return errors.New("unable to convert to int")
	}

	if len(includedCategory) != 0 && totalAmount != 0 {
		if savedSpent == 0 || savedRemaining == 0 {
			query := "UPDATE Budget SET spent=$1, remaining=$2 WHERE categories=$3"
			_, err = db.Exec(query, totalSpent, remainingBalance, category)
			if err != nil {
				return err
			}
		} else if savedRemaining != 0 && (spent <= savedRemaining || spent > savedRemaining) {
			query := "UPDATE Budget SET spent=$1, remaining=$2 WHERE categories=$3"
			_, err = db.Exec(query, totalSpent, savedRemaining-spent, category)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("first enter the total amount. see 'flow total-amount -h'")
	}
	return nil
}

func (m MyBudgetDatabase) GetBudgetData(filepath, filename string) error {
	bv := new(entities.BudgetVariables)
	db, err := m.Connection()
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

		var data []string
		amountStr := m.IntToString(bv.Amount)
		data = append(data, bv.Category, amountStr)
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
