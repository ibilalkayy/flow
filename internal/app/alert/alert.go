package internal_alert

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ibilalkayy/flow/db/alert_db"
	"github.com/ibilalkayy/flow/db/budget_db"
	internal_budget "github.com/ibilalkayy/flow/internal/app/budget"
	internal_spending "github.com/ibilalkayy/flow/internal/app/spend"
	"github.com/ibilalkayy/flow/internal/common/structs"
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

		if categoryAmount != 0 {
			err := alert_db.CreateAlert(av, "db/migrations/")
			if err != nil {
				return err
			}
			fmt.Printf("Alert is set for the '%s' category\n", av.Category)
		} else {
			return errors.New("category amount is not present")
		}
	} else {
		return errors.New("enter all the flags properly")
		// fmt.Printf("You can't spend more than your '%s' category budget\n", av.Category)
	}
	return nil
}

func SendAlert(category string) error {
	value, err := alert_db.ViewAlert(category)
	if err != nil {
		return err
	}

	method, ok1 := value[2].(string)
	frequency, ok2 := value[3].(string)
	day, ok3 := value[4].(int)
	weekdayStr, ok4 := value[5].(string)
	hour, ok5 := value[6].(int)
	minute, ok6 := value[7].(int)
	second, ok7 := value[8].(int)

	if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 || !ok6 || !ok7 {
		return errors.New("unable to convert string to int and string")
	}

	var weekday time.Weekday
	switch weekdayStr {
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

	switch method {
	case "email":
		switch frequency {
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
	details, err := budget_db.ViewBudget(category)
	if err != nil {
		return err
	}

	budgetAmount, ok1 := details[2].(int)
	spentAmount, ok2 := details[3].(int)

	if !ok1 || !ok2 {
		return errors.New("unable to convert spent or amount to int")
	}

	if spentAmount > budgetAmount {
		SendAlert(category)
	} else {
		fmt.Printf("The '%s' category amount is not exceeded", category)
	}
	return nil
}
