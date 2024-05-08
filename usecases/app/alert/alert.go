package usecases_alert

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ibilalkayy/flow/entities"
)

func (m MyAlerts) AlertSetup(av *entities.AlertVariables) error {
	if len(av.Category) != 0 && len(av.Frequency) != 0 && len(av.Method) != 0 && av.Days != 0 && len(av.Weekdays) != 0 && av.Hours != 0 && av.Minutes != 0 && (av.Seconds >= 0 && av.Seconds <= 60) {
		validMethods := map[string]bool{"email": true, "cli": true}
		validFrequencies := map[string]bool{"hourly": true, "daily": true, "weekly": true, "monthly": true}

		if !validMethods[strings.ToLower(av.Method)] {
			return errors.New("invalid alert method")
		}

		if !validFrequencies[strings.ToLower(av.Frequency)] {
			return errors.New("invalid alert frequency")
		}

		category, categoryAmount, err := m.CategoryAmount(av.Category)
		if err != nil {
			return err
		}

		if len(category) != 0 && categoryAmount != 0 {
			err := m.CreateAlert(av)
			if err != nil {
				return err
			}
			fmt.Printf("Alert is set for the '%s' category\n", av.Category)
		} else {
			return errors.New("budget data is not found. first set the budget")
		}
	} else {
		return errors.New("enter all the required flags properly")
	}
	return nil
}

func (m MyAlerts) SendAlert(category string) error {
	value, err := m.ViewAlert(category)
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
			m.HourlyNotification(category)
		case "daily":
			m.DailyNotification(hour, minute, second, category)
		case "weekly":
			m.WeeklyNotification(weekday, hour, minute, second, category)
		case "monthly":
			m.MonthlyNotification(day, hour, minute, second, category)
		default:
			return errors.New("wrong or no frequency is selected")
		}
	case "cli":
		return errors.New("you can't spend above your budget limit")
	}
	return nil
}

func (m MyAlerts) CheckNotification(category string) error {
	details, err := m.ViewBudget(category)
	if err != nil {
		return err
	}

	budgetAmount, ok1 := details[2].(int)
	spentAmount, ok2 := details[3].(int)

	if !ok1 || !ok2 {
		return errors.New("unable to convert spent or amount to int")
	}

	if spentAmount > budgetAmount {
		m.SendAlert(category)
	} else {
		fmt.Printf("The '%s' category amount is not exceeded\n", category)
	}
	return nil
}
