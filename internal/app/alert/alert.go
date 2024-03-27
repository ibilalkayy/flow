package internal_alert

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ibilalkayy/flow/db/alert_db"
	"github.com/ibilalkayy/flow/email"
	internal_budget "github.com/ibilalkayy/flow/internal/app/budget"
	"github.com/ibilalkayy/flow/internal/structs"
)

func CheckMethod(method, category string) error {
	if method == "email" {
		err := email.SendAlertEmail(category)
		if err != nil {
			return err
		}
	} else if method == "cli" {
		fmt.Println("cli is called")
	} else {
		return errors.New("write the correct method")
	}
	return nil
}

func AlertSetup(av *structs.AlertVariables) error {
	if len(av.Category) != 0 && len(av.Frequency) != 0 && len(av.Method) != 0 {
		validMethods := map[string]bool{"email": true, "cli": true}
		validFrequencies := map[string]bool{"hourly": true, "daily": true, "weekly": true, "monthly": true}

		if !validMethods[strings.ToLower(av.Method)] {
			return errors.New("invalid alert method")
		}

		if !validFrequencies[strings.ToLower(av.Frequency)] {
			return errors.New("invalid alert frequency")
		}

		categoryAmount, err := internal_budget.CategoryAmount(av.Category)
		if err != nil {
			return err
		}

		if len(categoryAmount) != 0 {
			err := alert_db.CreateAlert(av, "db/migrations/")
			if err != nil {
				return err
			}
			err = CheckMethod(av.Method, av.Category)
			if err != nil {
				return err
			}
			fmt.Println("Alert is set for a specific category")
		} else {
			return errors.New("category amount is not present")
		}
	} else {
		return errors.New("enter all the alert values")
	}
	return nil
}

// func AlertMessage() error {
// 	totalAmount, err := internal_budget.TotalBudgetAmount()
// 	if err != nil {
// 		return err
// 	}
// 	transactionAmount := transaction.TransactionAmount()

// 	if transactionAmount >= totalAmount {
// 		fmt.Printf("You can't spend more becuase your budget is set to %d\n", totalAmount)
// 	} else {
// 		return errors.New("enjoy your spending")
// 	}
// 	return nil
// }
