package internal_alert

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ibilalkayy/flow/db/alert_db"
	"github.com/ibilalkayy/flow/db/spend_db"
	internal_budget "github.com/ibilalkayy/flow/internal/app/budget"
	internal_spending "github.com/ibilalkayy/flow/internal/app/spend"
	"github.com/ibilalkayy/flow/internal/structs"
)

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
			fmt.Printf("Alert is set for the '%s' category", av.Category)
		} else {
			return errors.New("category amount is not present")
		}
	} else {
		fmt.Printf("You can't more than your '%s' category budget", av.Category)
	}
	return nil
}

func SendAlert(category string) error {
	value, err := alert_db.ViewAlert(category)
	if err != nil {
		return err
	}

	day, _ := strconv.Atoi(value[4])
	hour, _ := strconv.Atoi(value[6])
	minute, _ := strconv.Atoi(value[7])
	second, _ := strconv.Atoi(value[8])

	var weekday time.Weekday

	switch value[5] {
	case "monday":
		weekday = time.Monday
	case "tuesday":
		weekday = time.Tuesday
	case "wednesday":
		weekday = time.Wednesday
	case "thursday":
		weekday = time.Thursday
	case "friday":
		weekday = time.Friday
	case "saturday":
		weekday = time.Saturday
	case "sunday":
		weekday = time.Sunday
	default:
		return errors.New("wrong weekday is selected")
	}

	switch value[2] {
	case "email":
		switch value[3] {
		case "hourly":
			internal_spending.HourlyNotification(category)
		case "daily":
			internal_spending.DailyNotification(hour, minute, second, category)
		case "weekly":
			internal_spending.WeeklyNotification(weekday, hour, minute, second, category)
		case "monthly":
			internal_spending.MonthlyNotification(day, hour, minute, second, category)
		default:
			return errors.New("wrong or no frequency is selected")
		}
	case "cli":
		return errors.New("you can't spend above your budget limit")
	}
	return nil
}

func CheckNotification(category string) error {
	value, err := spend_db.ViewSpending(category)
	if err != nil {
		return err
	}

	if value[3] == "true" {
		SendAlert(category)
	} else {
		fmt.Printf("The '%s' category amount is not exceeded", category)
	}
	return nil
}
