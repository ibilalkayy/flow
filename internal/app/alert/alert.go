package internal_alert

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/ibilalkayy/flow/cmd/transaction"
	"github.com/ibilalkayy/flow/db/budget_db"
	internal_budget "github.com/ibilalkayy/flow/internal/app/budget"
	"github.com/ibilalkayy/flow/internal/structs"
)

func CreateAlert(av *structs.AlertVariables, basePath string) error {
	data, err := budget_db.Table(basePath, "002_create_alert_table.sql", 0)
	if err != nil {
		return err
	}

	query := "INSERT INTO Alert(alert_methods, alert_frequencies) VALUES($1, $2)"
	insert, err := data.Prepare(query)
	if err != nil {
		return err
	}
	defer insert.Close()

	if len(av.Method) != 0 && len(av.Frequency) != 0 {
		_, err = insert.Exec(av.Method, av.Frequency)
		if err != nil {
			return err
		}
		fmt.Println("Alert data is successfully inserted!")
	} else {
		return errors.New("enter both the method and frequency")
	}
	return nil
}

func Notification(av *structs.AlertVariables) error {
	if av.Method == "email" || av.Method == "Email" || av.Method == "cli" || av.Method == "Cli" || av.Method == "CLI" {
		if av.Frequency == "hourly" || av.Frequency == "Hourly" || av.Frequency == "daily" || av.Frequency == "Daily" || av.Frequency == "weekly" || av.Frequency == "Weekly" || av.Frequency == "monthly" || av.Frequency == "Monthly" {
			err := CreateAlert(av, "db/budget_db/migrations/")
			if err != nil {
				return err
			}
		} else {
			return errors.New("invalid alert frequency")
		}
	} else {
		return errors.New("invalid alert method")
	}
	return nil
}

func AlertSetup(av structs.AlertVariables) error {
	if len(av.Frequency) != 0 && len(av.Method) != 0 {
		if len(av.Total) != 0 {
			budgetAmount, err := internal_budget.TotalBudgetAmount()
			if err != nil {
				return err
			}
			totalAmount := strconv.Itoa(budgetAmount)

			if len(totalAmount) != 0 && av.Total == "amount" {
				fmt.Println("Alert is set for the total amount")
			} else {
				return errors.New("total amount is not present or the flag value is not given properly")
			}
		} else if len(av.Category) != 0 {
			categoryAmount, err := internal_budget.CategoryAmount(av.Category)
			if err != nil {
				return err
			}

			if len(categoryAmount) != 0 {
				fmt.Println("Alert is set for a specific category")
			} else {
				return errors.New("category amount is not present")
			}
		} else {
			return errors.New("select a category")
		}

		err := Notification(&av)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Your alert information is setup successfully")
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
