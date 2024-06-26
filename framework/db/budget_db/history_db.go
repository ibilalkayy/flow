package budget_db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/entities"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (h MyBudgetDB) InsertHistory(hv *entities.HistoryVariables) error {
	data, err := h.Deps.Connect.Table("framework/db/migrations/001_create_budget_table.sql", 1)
	if err != nil {
		return err
	}

	query := "INSERT INTO History(dates, timez, categories, amounts, blockchains, addresses) VALUES($1, $2, $3, $4, $5, $6)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	includedCategory, value, err := h.Deps.TotalAmount.TotalAmountValues()
	if err != nil {
		return err
	}

	totalAmount, ok := value[0].(int)
	if !ok {
		return errors.New("unable to convert to int")
	}

	if len(hv.Category) != 0 && len(includedCategory) != 0 {
		if hv.Amount != 0 && totalAmount != 0 {
			_, err = insert.Exec(hv.Date, hv.Time, hv.Category, hv.Amount, hv.Blockchain, hv.RecipientAddress)
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

func (h MyBudgetDB) ViewHistory(category string) ([2]interface{}, error) {
	hv := new(entities.HistoryVariables)

	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return [2]interface{}{}, err
	}

	defer db.Close()

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Date", "Time", "Category", "Amounts", "Blockchains", "Addresses"})

	var rows *sql.Rows
	if len(category) != 0 {
		query := "SELECT dates, timez, categories, amounts, blockchains, addresses from History WHERE categories=$1"
		rows, err = db.Query(query, category)
	} else {
		query := "SELECT dates, timez, categories, amounts, blockchains, addresses from History"
		rows, err = db.Query(query)
	}
	if err != nil {
		return [2]interface{}{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&hv.Date, &hv.Time, &hv.Category, &hv.Amount, &hv.Blockchain, &hv.RecipientAddress); err != nil {
			return [2]interface{}{}, err
		}

		if len(hv.Category) != 0 && hv.Amount != 0 {
			tw.AppendRow([]interface{}{hv.Date, hv.Time, hv.Category, hv.Amount, hv.Blockchain, hv.RecipientAddress})
		}
	}

	tableRender := "History Data\n" + tw.Render()
	details := [2]interface{}{tableRender, hv.Category}
	return details, nil
}

func (h MyBudgetDB) RemoveHistory(category string) error {
	db, err := h.Deps.Connect.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	data, err := h.ViewHistory(category)
	if err != nil {
		return err
	}

	foundCategory, ok := data[1].(string)
	if !ok {
		return errors.New("unable to convert data to string")
	}

	query := "DELETE FROM History"
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
			fmt.Println("History data is successfully deleted!")
		} else {
			return errors.New("no data is found")
		}
	}

	return nil
}
