package internal_spending

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ibilalkayy/flow/db/alert_db"
)

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
			HourlyNotification(category)
		case "daily":
			DailyNotification(hour, minute, second, category)
		case "weekly":
			WeeklyNotification(weekday, hour, minute, second, category)
		case "monthly":
			MonthlyNotification(day, hour, minute, second, category)
		default:
			return errors.New("wrong or no frequency is selected")
		}
	case "cli":
		return errors.New("you can't spend above your budget limit")
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
			if err := SendAlert(category); err != nil {
				return err
			}
		}
	} else {
		return errors.New("category is not found. first setup the alert. see 'flow budget alert setup -h'")
	}
	return nil
}
