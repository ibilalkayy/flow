package total_amount_db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/db"
	"github.com/ibilalkayy/flow/internal/common/structs"
	"github.com/jedib0t/go-pretty/v6/table"
)

func SetTotalAmount(tv *structs.TotalAmountVariables, basepath string) error {
	data, err := db.Table(basepath, "003_create_total_amount_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO TotalAmount(total_amount, remaining_amount, included_category, label, statuss) VALUES($1, $2, $3, $4, $5)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}

	defer insert.Close()

	if tv.TotalAmount != 0 && len(tv.Included) != 0 {
		_, err = insert.Exec(tv.TotalAmount, tv.RemainingAmount, tv.Included, tv.Label, tv.Status)
		if err != nil {
			return err
		}
		fmt.Println("Total amount data is successfully inserted!")
	} else {
		return errors.New("write total amount and category. see 'flow total-amount set -h'")
	}
	return nil
}

func ViewTotalAmount() ([4]interface{}, error) {
	tv := new(structs.TotalAmountVariables)

	db, err := db.Connection()
	if err != nil {
		return [4]interface{}{}, err
	}
	defer db.Close()

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Total Amount", "Remaining Amount", "Included Category", "Label", "Status"})

	var rows *sql.Rows
	query := "SELECT total_amount, remaining_amount, included_category, label, statuss FROM TotalAmount"
	rows, err = db.Query(query)
	if err != nil {
		return [4]interface{}{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&tv.TotalAmount, &tv.RemainingAmount, &tv.Included, &tv.Label, &tv.Status); err != nil {
			return [4]interface{}{}, nil
		}
	}
	// Append data to the table inside the loop
	tw.AppendRow([]interface{}{tv.TotalAmount, tv.RemainingAmount, tv.Included, tv.Label, tv.Status})
	tableRender := "Total Amount\n" + tw.Render()
	details := [4]interface{}{tableRender, tv.Included, tv.TotalAmount, tv.Status}
	return details, nil
}

func RemoveTotalAmount() error {
	db, err := db.Connection()
	if err != nil {
		return err
	}

	query := "DELETE FROM TotalAmount"
	remove, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer remove.Close()

	_, err = remove.Exec()
	if err != nil {
		return nil
	}

	fmt.Println("Tottal amount is successfully removed!")
	return nil
}

func UpdateTotalAmount(tv *structs.TotalAmountVariables) error {
	var query string
	var params []interface{}

	db, err := db.Connection()
	if err != nil {
		return err
	}
	if tv.TotalAmount != 0 && len(tv.Label) != 0 {
		query = "UPDATE TotalAmount SET total_amount=$1, label=$2"
		params = []interface{}{tv.TotalAmount, tv.Label}
	} else if tv.TotalAmount != 0 {
		query = "UPDATE TotalAmount SET total_amount=$1"
		params = []interface{}{tv.TotalAmount}
	} else if len(tv.Label) != 0 {
		query = "UPDATE TotalAmount SET label=$1"
		params = []interface{}{tv.Label}
	} else {
		return errors.New("no flag is provided to update")
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		return err
	}
	fmt.Println("Total amount data is successfully updated!")
	return nil
}

func UpdateStatus(tv *structs.TotalAmountVariables) error {
	db, err := db.Connection()
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

func CalculateRemaining(category string) error {
	db, err := db.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Find the total amount data
	var totalAmount int
	if len(category) != 0 {
		query := "SELECT total_amount FROM TotalAmount WHERE included_category=$1"
		err := db.QueryRow(query, category).Scan(&totalAmount)
		if err != nil {
			return err
		}
	} else {
		return errors.New("category is not present")
	}

	// Find the budget amount data
	var savedSpent int
	if len(category) != 0 {
		query := "SELECT spent FROM Budget WHERE categories = $1"
		err := db.QueryRow(query, category).Scan(&savedSpent)
		if err != nil {
			return err
		}
	} else {
		return errors.New("category is not present")
	}

	remainingBalance := totalAmount - savedSpent
	query := "UPDATE TotalAmount SET remaining_amount=$1 WHERE included_category=$2"
	_, err = db.Exec(query, remainingBalance, category)
	if err != nil {
		return err
	}
	return nil
}
