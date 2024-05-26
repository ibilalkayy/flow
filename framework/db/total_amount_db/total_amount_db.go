package total_amount_db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/entities"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (h MyTotalAmountDB) InsertTotalAmount(tv *entities.TotalAmountVariables) error {
	data, err := h.Deps.Connect.Table("framework/db/migrations/003_create_total_amount_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO TotalAmount(total_amount, spent_amount, remaining_amount, statuss) VALUES($1, $2, $3, $4)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}

	defer insert.Close()

	if tv.TotalAmount != 0 {
		_, err = insert.Exec(tv.TotalAmount, tv.SpentAmount, tv.RemainingAmount, tv.Status)
		if err != nil {
			return err
		}
		fmt.Println("Total amount data is successfully inserted!")
	} else {
		return errors.New("write total amount and category. see 'flow total-amount set -h'")
	}
	return nil
}

func (h MyTotalAmountDB) ViewTotalAmount() ([5]interface{}, error) {
	tv := new(entities.TotalAmountVariables)

	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return [5]interface{}{}, err
	}
	defer db.Close()

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Total Amount", "Spent Amount", "Remaining Amount", "Status"})

	var rows *sql.Rows
	query := "SELECT total_amount, spent_amount, remaining_amount, statuss FROM TotalAmount"
	rows, err = db.Query(query)
	if err != nil {
		return [5]interface{}{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&tv.TotalAmount, &tv.SpentAmount, &tv.RemainingAmount, &tv.Status); err != nil {
			return [5]interface{}{}, err
		}
	}

	// Append data to the table inside the loop
	tw.AppendRow([]interface{}{tv.TotalAmount, tv.SpentAmount, tv.RemainingAmount, tv.Status})
	tableRender := "Total Amount\n" + tw.Render()
	details := [5]interface{}{tableRender, tv.TotalAmount, tv.SpentAmount, tv.RemainingAmount, tv.Status}
	return details, nil
}

func (h MyTotalAmountDB) RemoveTotalAmount(category string) error {
	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	if len(category) != 0 {
		query := "DELETE FROM TotalAmountCategories WHERE included_categories=$1"
		removeCategory, err := db.Prepare(query)
		if err != nil {
			return err
		}
		defer removeCategory.Close()

		_, err = removeCategory.Exec(category)
		if err != nil {
			return err
		}

		query = "DELETE FROM Budget WHERE categories=$1"
		removeBudgetCategory, err := db.Prepare(query)
		if err != nil {
			return err
		}
		defer removeBudgetCategory.Close()

		_, err = removeBudgetCategory.Exec(category)
		if err != nil {
			return err
		}
		fmt.Printf("'%s' category removed successfully from Total Amount and Budget!\n", category)
	} else {
		query := "DELETE FROM TotalAmountCategories"
		removeTotalAmountCategories, err := db.Prepare(query)
		if err != nil {
			return err
		}
		defer removeTotalAmountCategories.Close()

		_, err = removeTotalAmountCategories.Exec()
		if err != nil {
			return err
		}

		query = "DELETE FROM Budget"
		removeBudget, err := db.Prepare(query)
		if err != nil {
			return err
		}
		defer removeBudget.Close()

		_, err = removeBudget.Exec()
		if err != nil {
			return err
		}
	}

	if len(category) == 0 {
		query := "DELETE FROM TotalAmount"
		removeTotalAmount, err := db.Prepare(query)
		if err != nil {
			return err
		}
		defer removeTotalAmount.Close()

		_, err = removeTotalAmount.Exec()
		if err != nil {
			return err
		}
		fmt.Println("All entries removed successfully from TotalAmount and Budget!")
	}

	return nil
}

func (h MyTotalAmountDB) UpdateTotalAmount(tv *entities.TotalAmountVariables) error {
	var query string
	var params []interface{}

	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return err
	}

	details, err := h.Deps.TotalAmount.ViewTotalAmount()
	if err != nil {
		return err
	}

	totalAmountInDB, ok1 := details[1].(int)
	spentAmountInDB, ok2 := details[2].(int)
	remainingAmountInDB, ok3 := details[3].(int)
	if !ok1 || !ok2 || !ok3 {
		return errors.New("unable to convert amount to int")
	}

	if tv.TotalAmount != 0 {
		if tv.TotalAmount > totalAmountInDB {
			// Update Remaining Amount if Total Amount increases
			updatedRemaining := tv.TotalAmount - totalAmountInDB
			remainingAmountInDB += updatedRemaining
		} else if tv.TotalAmount < totalAmountInDB {
			// Update Spent Amount and Remaining Amount if Total Amount decreases
			if spentAmountInDB <= tv.TotalAmount {
				remainingAmountInDB = tv.TotalAmount - spentAmountInDB
			} else {
				spentAmountInDB = tv.TotalAmount
				remainingAmountInDB = 0
			}
		} else {
			return errors.New("this amount is already present. enter a different amount")
		}
		query = "UPDATE TotalAmount SET total_amount = $1, spent_amount = $2, remaining_amount = $3"
		params = []interface{}{tv.TotalAmount, spentAmountInDB, remainingAmountInDB}
	} else if len(tv.Included) != 0 && len(tv.NewCategory) != 0 {
		query = "UPDATE TotalAmountCategories SET included_categories=$1 WHERE included_categories=$2"
		params = []interface{}{tv.NewCategory, tv.Included}
	} else if len(tv.Label) != 0 && len(tv.Included) != 0 {
		query = "UPDATE TotalAmountCategories SET labels=$1 WHERE included_categories=$2"
		params = []interface{}{tv.Label, tv.Included}
	} else {
		return errors.New("write the present category also to update the values")
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		return err
	}
	fmt.Println("Total amount data is successfully updated!")
	return nil
}

func (h MyTotalAmountDB) UpdateStatus(tv *entities.TotalAmountVariables) error {
	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return err
	}

	if len(tv.Status) != 0 && tv.Status == "Active" {
		query := "UPDATE TotalAmount SET statuss=$1"
		_, err = db.Exec(query, "active")
		if err != nil {
			return err
		}
		fmt.Println("Total amount is actived")
	} else {
		query := "UPDATE TotalAmount SET statuss=$1"
		_, err = db.Exec(query, "inactive")
		if err != nil {
			return err
		}
		fmt.Println("Total amount is inactived")
	}
	return nil
}

func (h MyTotalAmountDB) CalculateRemaining(category string) error {
	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Find the total amount data
	var totalAmount int
	if len(category) != 0 {
		query := "SELECT total_amount FROM TotalAmount"
		err := db.QueryRow(query).Scan(&totalAmount)
		if err != nil {
			return err
		}
	} else {
		return errors.New("category is not found")
	}

	// Calculate the total spent amount from Budget
	query := "SELECT SUM(spent) FROM Budget"
	var totalSpent sql.NullInt64
	err = db.QueryRow(query).Scan(&totalSpent)
	if err != nil {
		return err
	}

	// Check if totalSpent is valid, if not set it to zero
	var savedSpent int
	if totalSpent.Valid {
		savedSpent = int(totalSpent.Int64)
	} else {
		savedSpent = 0
	}

	// Calculate remaining balance
	remainingBalance := totalAmount - savedSpent

	// Update the remaining amount and spent amount in TotalAmount table
	updateQuery := "UPDATE TotalAmount SET spent_amount=$1, remaining_amount=$2"
	_, err = db.Exec(updateQuery, savedSpent, remainingBalance)
	if err != nil {
		return err
	}

	return nil
}
