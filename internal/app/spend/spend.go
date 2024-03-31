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

func HourlyNotification(category string) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		<-ticker.C
		email.SendAlertEmail(category)
		fmt.Println("Printed every hour")
	}
}

func DailyNotification(hour, min, sec int, category string) {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, now.Location())
	if next.Before(now) {
		next = next.Add(24 * time.Hour)
	}
	fmt.Printf("Next daily print will be at %s\n", next)
	time.Sleep(next.Sub(now))
	fmt.Println("Printed daily at the specified time")
	email.SendAlertEmail(category)
}

func WeeklyNotification(weekday time.Weekday, hour, min, sec int, category string) {
	now := time.Now()
	daysUntilNextWeekday := int((weekday - now.Weekday() + 7) % 7)
	next := time.Date(now.Year(), now.Month(), now.Day()+daysUntilNextWeekday, hour, min, sec, 0, now.Location())
	fmt.Printf("Next weekly print will be on %s at %s\n", weekday, next)
	time.Sleep(next.Sub(now))
	fmt.Println("Printed weekly on the specified day and time")
	email.SendAlertEmail(category)
}

func MonthlyNotification(day, hour, min, sec int, category string) {
	now := time.Now()
	year, month, _ := now.Date()
	next := time.Date(year, month, day, hour, min, sec, 0, now.Location())
	if next.Before(now) {
		next = next.AddDate(0, 1, 0)
	}
	fmt.Printf("Next monthly print will be on the %dth day at %s\n", day, next)
	time.Sleep(next.Sub(now))
	fmt.Println("Printed monthly on the specified day and time")
	email.SendAlertEmail(category)
}
