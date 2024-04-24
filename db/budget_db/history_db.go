package budget_db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/db"
	"github.com/ibilalkayy/flow/internal/common/functions"
	"github.com/ibilalkayy/flow/internal/common/structs"
	"github.com/jedib0t/go-pretty/v6/table"
)

func InsertHistory(hv *structs.HistoryVariables, basePath string) error {
	data, err := db.Table(basePath, "001_create_budget_table.sql", 1)
	if err != nil {
		return err
	}

	query := "INSERT INTO History(dates, timez, categories, amounts, transaction_ids, blockchains, addresses) VALUES($1, $2, $3, $4, $5, $6, $7)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	includedCategory, value, err := functions.TotalAmountValues()
	if err != nil {
		return err
	}

	totalAmount, ok := value[0].(int)
	if !ok {
		return errors.New("unable to convert to int")
	}

	if len(hv.Category) != 0 && len(includedCategory) != 0 {
		if hv.Amount != 0 && totalAmount != 0 {
			_, err = insert.Exec(hv.Date, hv.Time, hv.Category, hv.Amount, hv.TransactionID, hv.Blockchain, hv.Address)
			if err != nil {
				return err
			}
		} else {
			return errors.New("enter the amount")
		}
	} else {
		return errors.New("enter the category")
	}
	return nil
}

func ViewHistory(category string) error {
	hv := new(structs.HistoryVariables)

	db, err := db.Connection()
	if err != nil {
		return err
	}

	defer db.Close()

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Date", "Time", "Category", "Amounts", "Transaction IDs", "Blockchains", "Addresses"})

	var rows *sql.Rows
	if len(category) != 0 {
		query := "SELECT dates, timez, categories, amounts, transaction_ids, blockchains, addresses from History WHERE categories=$1"
		rows, err = db.Query(query, category)
	} else {
		query := "SELECT dates, timez, categories, amounts, transaction_ids, blockchains, addresses from History"
		rows, err = db.Query(query)
	}
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&hv.Date, &hv.Time, &hv.Category, &hv.Amount, &hv.TransactionID, &hv.Blockchain, &hv.Address); err != nil {
			return err
		}

		if len(hv.Category) != 0 && hv.Amount != 0 {
			tw.AppendRow([]interface{}{hv.Date, hv.Time, hv.Category, hv.Amount, hv.TransactionID, hv.Blockchain, hv.Address})
		}
	}

	tableRender := "History Data\n" + tw.Render()
	fmt.Println(tableRender)
	return nil
}

func RemoveHistory(category string) error {
	db, err := db.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM History"
	var args []interface{}

	if len(category) != 0 {
		query += " WHERE categories=$1"
		args = append(args, category)
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
		fmt.Printf("History is successfully deleted!")
	}

	return nil
}
