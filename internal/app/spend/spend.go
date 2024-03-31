package internal_spending

import (
	"errors"
	"fmt"
	"time"

	"github.com/ibilalkayy/flow/db/alert_db"
	"github.com/ibilalkayy/flow/email"
)

func AlertMethod(category string) error {
	method, err := alert_db.ViewAlert(category)
	if err != nil {
		return err
	}

	if method[2] == "email" {
		err := email.SendAlertEmail(category)
		if err != nil {
			return err
		}
	} else if method[2] == "cli" {
		fmt.Println("you can't spend above your budget limit")
	} else {
		return errors.New("write the correct method")
	}
	return nil
}

func AlertFrequency(category string) error {
	frequency, err := alert_db.ViewAlert(category)
	if err != nil {
		return err
	}

	if frequency[3] == "hourly" {
		HourlyNotification(category)
	} else if frequency[3] == "daily" {
		DailyNotification(11, 37, 0, category)
	} else if frequency[3] == "weekly" {
		WeeklyNotification(time.Sunday, 11, 53, 0, category)
	} else if frequency[3] == "monthly" {
		MonthlyNotification(31, 12, 19, 0, category)
	} else {
		return errors.New("select the right frequency")
	}
	return nil
}

func GiveAlerts(category string) error {
	err := AlertMethod(category)
	if err != nil {
		return err
	}
	err = AlertFrequency(category)
	if err != nil {
		return err
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
			if err := GiveAlerts(category); err != nil {
				return err
			}
		}
	} else {
		return errors.New("category is not found. first setup the alert. see 'flow budget alert setup -h'")
	}
	return nil
}
