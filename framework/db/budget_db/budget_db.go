package budget_db

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/handler"
	"github.com/jedib0t/go-pretty/v6/table"
)

type MyBudgetDB struct {
	*handler.Handler
}

func (h MyBudgetDB) CreateBudget(bv *entities.BudgetVariables) error {
	data, err := h.Deps.Connect.Table("framework/db/migrations/001_create_budget_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO Budget(categories, amounts, spent, remaining) VALUES($1, $2, $3, $4)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	includedCategory, value, err := h.Deps.TotalAmount.TotalAmountValues()
	if err != nil {
		return err
	}

	fullTotalAmount, ok := value[0].(int)
	if !ok {
		return errors.New("unable to convert to int")
	}

	budgetAmount, err := h.Deps.ManageBudget.TakeBudgetAmount()
	if err != nil {
		return nil
	}

	totalBudgetAmount := 0
	for _, amount := range budgetAmount {
		totalBudgetAmount += amount
	}

	for i := 0; i < len(includedCategory); i++ {
		if len(bv.Category) != 0 && len(includedCategory) != 0 && includedCategory[i][0] == bv.Category && fullTotalAmount != 0 && totalBudgetAmount+bv.Amount <= fullTotalAmount {
			_, err = insert.Exec(bv.Category, bv.Amount, 0, bv.Amount)
			if err != nil {
				return err
			}
			fmt.Println("Budget data is successfully inserted!")
			return nil
		}
	}
	return errors.New("enter the category in the total amount or set the budget below or equal to total amount. see 'flow total-amount -h'")
}

func (h MyBudgetDB) ViewBudget(category string) ([5]interface{}, error) {
	// Create a new instance of BudgetVariables to hold the retrieved data
	bv := new(entities.BudgetVariables)

	// Connect to the database
	db, err := h.Deps.Connect.Connection()
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

func (h MyBudgetDB) RemoveBudget(category string) error {
	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	data, err := h.ViewBudget(category)
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
			fmt.Println("Budget data is successfully deleted!")
		} else {
			return errors.New("no data is found")
		}
	}

	return nil
}

func (h MyBudgetDB) UpdateBudget(bv *entities.BudgetVariables, new_category string) error {
	var count int
	var query string
	var params []interface{}

	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return err
	}

	// Check if the old category exists
	err = db.QueryRow("SELECT COUNT(*) FROM Budget WHERE categories = $1", bv.Category).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("'" + bv.Category + "'" + " category does not exist")
	}

	includedCategory, value, err := h.Deps.TotalAmount.TotalAmountValues()
	if err != nil {
		return err
	}

	fullTotalAmount, ok := value[0].(int)
	if !ok {
		return errors.New("unable to convert to int")
	}

	budgetAmount, err := h.Deps.ManageBudget.TakeBudgetAmount()
	if err != nil {
		return nil
	}

	totalBudgetAmount := 0
	for _, amount := range budgetAmount {
		totalBudgetAmount += amount
	}

	expectional_budget_amount, err := h.Deps.ManageBudget.BudgetAmountWithException(bv)
	if err != nil {
		return err
	}

	details, err := h.Deps.ManageBudget.ViewBudget(bv.Category)
	if err != nil {
		return err
	}

	budgetAmountInDB, ok1 := details[2].(int)
	spentAmountInDB, ok2 := details[3].(int)
	remainingAmountInDB, ok3 := details[4].(int)
	if !ok1 || !ok2 || !ok3 {
		return errors.New("unable to convert to int")
	}

	cr := entities.BudgetCalculateVariables{
		BudgetAmount:        bv.Amount,
		BudgetAmountInDB:    budgetAmountInDB,
		SpentAmountInDB:     spentAmountInDB,
		RemainingAmountInDB: remainingAmountInDB,
	}

	if len(includedCategory) != 0 && fullTotalAmount != 0 && bv.Amount <= fullTotalAmount && totalBudgetAmount <= fullTotalAmount && expectional_budget_amount+bv.Amount <= fullTotalAmount {
		if bv.Amount != 0 {
			data, err := h.Deps.ManageBudget.CalculateRemaining(&cr)
			if err != nil {
				return err
			}

			if len(new_category) != 0 {
				query = "UPDATE Budget SET categories=$1, amounts=$2, spent=$3, remaining=$4 WHERE categories=$5"
				params = []interface{}{new_category, bv.Amount, data[0], data[1], bv.Category}
			} else {
				query = "UPDATE Budget SET amounts=$1, spent=$2, remaining=$3 WHERE categories=$4"
				params = []interface{}{bv.Amount, data[0], data[1], bv.Category}
			}
		} else if len(new_category) != 0 {
			query = "UPDATE Budget SET categories=$1 WHERE categories=$2"
			params = []interface{}{new_category, bv.Category}
		} else {
			fmt.Println("No field provided to update")
		}

		_, err = db.Exec(query, params...)
		if err != nil {
			return err
		}
	} else {
		return errors.New("write the budget amount below or equal to the total amount. see 'flow total-amount -h'")
	}
	return nil
}

func (h MyBudgetDB) AddBudgetExpenditure(spent int, category string) error {
	db, err := h.Deps.Connect.Connection()
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

	includedCategory, value, err := h.Deps.TotalAmount.TotalAmountValues()
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

func (h MyBudgetDB) GetBudgetData(filepath, filename string) error {
	bv := new(entities.BudgetVariables)
	db, err := h.Deps.Connect.Connection()
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
		amountStr := h.Deps.Common.IntToString(bv.Amount)
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
