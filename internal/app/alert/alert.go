package internal_alert

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/ibilalkayy/flow/cmd/transaction"
	"github.com/ibilalkayy/flow/db/budget_db"
	internal_budget "github.com/ibilalkayy/flow/internal/app/budget"
	"github.com/ibilalkayy/flow/internal/structs"
)

func GenerateUniqueCategory() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return fmt.Sprintf("total_category_%s", b)
}

func CreateAlert(av *structs.AlertVariables, basePath string) error {
	data, err := budget_db.Table(basePath, "002_create_alert_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO Alert(categories, category_amounts, total_amount, alert_methods, alert_frequencies) VALUES($1, $2, $3, $4, $5)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	var total, category, categoryAmount string

	if len(av.Category) == 0 {
		category = GenerateUniqueCategory()
	} else {
		category = av.Category
	}

	if len(av.Total) != 0 && len(av.Method) != 0 && len(av.Frequency) != 0 {
		total = av.Total
	} else if len(av.Category) != 0 && len(av.Method) != 0 && len(av.Frequency) != 0 {
		category = av.Category
		categoryAmount, err = internal_budget.CategoryAmount(category)
		if err != nil {
			return err
		}
	} else {
		return errors.New("enter the required flags")
	}

	_, err = insert.Exec(category, categoryAmount, total, av.Method, av.Frequency)
	if err != nil {
		return err
	}
	return nil
}

func AlertSetup(av *structs.AlertVariables) error {
	if len(av.Frequency) != 0 && len(av.Method) != 0 {
		validMethods := map[string]bool{"email": true, "cli": true}
		validFrequencies := map[string]bool{"hourly": true, "daily": true, "weekly": true, "monthly": true}

		if !validMethods[strings.ToLower(av.Method)] {
			return errors.New("invalid alert method")
		}

		if !validFrequencies[strings.ToLower(av.Frequency)] {
			return errors.New("invalid alert frequency")
		}

		if len(av.Total) != 0 {
			budgetAmount, err := internal_budget.TotalBudgetAmount()
			if err != nil {
				return err
			}
			totalAmount := strconv.Itoa(budgetAmount)

			if len(totalAmount) != 0 && av.Total == totalAmount {
				err := CreateAlert(av, "db/budget_db/migrations/")
				if err != nil {
					return err
				}
				fmt.Println("Alert is set for the total amount")
			} else {
				return errors.New("total amount is not given. type 'flow budget view' to get the total amount")
			}
		} else if len(av.Category) != 0 {
			categoryAmount, err := internal_budget.CategoryAmount(av.Category)
			if err != nil {
				return err
			}

			if len(categoryAmount) != 0 {
				err := CreateAlert(av, "db/budget_db/migrations/")
				if err != nil {
					return err
				}
				fmt.Println("Alert is set for a specific category")
			} else {
				return errors.New("category amount is not present")
			}
		} else {
			return errors.New("select a category")
		}
	} else {
		return errors.New("enter all the alert values")
	}
	return nil
}

func AlertMessage() error {
	totalAmount, err := internal_budget.TotalBudgetAmount()
	if err != nil {
		return err
	}
	transactionAmount := transaction.TransactionAmount()

	if transactionAmount >= totalAmount {
		fmt.Printf("You can't spend more becuase your budget is set to %d\n", totalAmount)
	} else {
		return errors.New("enjoy your spending")
	}
	return nil
}
