package app

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/db"
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

func ViewBudget(category string) ([2]string, error) {
	bv := new(BudgetVariables)
	db, err := db.Connection()
	if err != nil {
		return [2]string{}, err
	}

	query := "SELECT categories, amounts FROM Budget WHERE categories=$1"
	if err := db.QueryRow(query, category).Scan(&bv.Category, &bv.Amount); err != nil {
		return [2]string{}, err
	}
	return [2]string{bv.Category, bv.Amount}, err
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
	db, err := db.Connection()
	if err != nil {
		return err
	}

	// Check if the old category exists
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Budget WHERE categories = $1", old).Scan(&count)
	if err != nil {
		return err
	}

	// If the old category does not exist, return an error
	if count == 0 {
		return errors.New("'" + old + "'" + " category does not exist")
	}

	// If the old category exists, execute the update query
	query := "UPDATE Budget SET categories=$1, amounts=$2 WHERE categories=$3"
	_, err = db.Exec(query, new, amount, old)
	if err != nil {
		return err
	}
	fmt.Println("Your budget category is successfully updated!")
	return nil
}
