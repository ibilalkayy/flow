package internal_spending

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/db/alert_db"
	"github.com/ibilalkayy/flow/email"
)

func AlertMethod(category string) error {
	values, err := alert_db.ViewAlert(category)
	if err != nil {
		return err
	}

	if values[2] == "email" {
		err := email.SendAlertEmail(category)
		if err != nil {
			return err
		}
	} else if values[2] == "cli" {
		fmt.Println("you can't spend above your budget limit")
	} else {
		return errors.New("write the correct method")
	}
	return nil
}

func SpendMoney(category, spending_amount string) error {
	values, err := alert_db.ViewAlert(category)
	if err != nil {
		return err
	}

	if category == values[0] {
		if spending_amount <= values[1] {
			fmt.Println("enjoy your spending")
		} else {
			err := AlertMethod(category)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("category is not found. first setup the alert. see 'flow budget alert setup -h'")
	}
	return nil
}
