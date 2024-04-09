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

	query := "INSERT INTO TotalAmount(amount, included_category, excluded_category, label) VALUES($1, $2, $3, $4)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}

	defer insert.Close()

	if tv.Amount != 0 {
		_, err = insert.Exec(tv.Amount, tv.Included, tv.Excluded, tv.Label)
		if err != nil {
			return err
		}
		fmt.Println("Total amount data is successfully inserted!")
	} else {
		return errors.New("total amount can't be empty")
	}
	return nil
}

func ViewTotalAmount() ([2]interface{}, error) {
	tv := new(structs.TotalAmountVariables)

	db, err := db.Connection()
	if err != nil {
		return [2]interface{}{}, err
	}
	defer db.Close()

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Total Amount", "Included Category", "Excluded Category", "Label"})

	var rows *sql.Rows
	query := "SELECT amount, included_category, excluded_category, label FROM TotalAmount"
	rows, err = db.Query(query)
	if err != nil {
		return [2]interface{}{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&tv.Amount, &tv.Included, &tv.Excluded, &tv.Label); err != nil {
			return [2]interface{}{}, nil
		}
	}
	// Append data to the table inside the loop
	tw.AppendRow([]interface{}{tv.Amount, tv.Included, tv.Excluded, tv.Label})
	tableRender := "Total Amount\n" + tw.Render()
	details := [2]interface{}{tableRender, tv.Amount}
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

	if tv.Amount != 0 {
		query = "UPDATE TotalAmount SET amount=$1"
		params = []interface{}{tv.Amount}
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
