package usecases_alert

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/handler"
)

type MyAlert struct {
	*handler.Handler
}

func (h MyAlert) AlertSetup(av *entities.AlertVariables) error {
	if len(av.Category) != 0 && len(av.Frequency) != 0 && len(av.Method) != 0 {
		validMethods := map[string]bool{"email": true, "cli": true}
		validFrequencies := map[string]bool{"hourly": true, "daily": true, "weekly": true, "monthly": true}

		if !validMethods[strings.ToLower(av.Method)] {
			return errors.New("invalid alert method")
		}

		if !validFrequencies[strings.ToLower(av.Frequency)] {
			return errors.New("invalid alert frequency")
		}

		category, categoryAmount, err := h.Deps.Budget.CategoryAmount(av.Category)
		if err != nil {
			return err
		}

		if len(category) != 0 && categoryAmount != 0 {
			err := h.Deps.AlertDB.CreateAlert(av)
			if err != nil {
				return err
			}
			fmt.Printf("Alert is set for the '%s' category\n", av.Category)
		} else {
			return errors.New("first create a budget. go to 'flow budget -h' for help")
		}
	} else {
		return errors.New("enter all the required flags properly")
	}
	return nil
}

func (h MyAlert) SendAlert(category string) error {
	value, err := h.Deps.AlertDB.ViewAlert(category)
	if err != nil {
		return err
	}

	method, ok1 := value[2].(string)
	frequency, ok2 := value[3].(string)
	day, ok3 := value[4].(int)
	weekdayStr, ok4 := value[5].(string)
	hour, ok5 := value[6].(int)
	minute, ok6 := value[7].(int)

	if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 || !ok6 {
		return errors.New("unable to convert string to int and string")
	}

	weekdayStr = strings.TrimSpace(strings.ToLower(weekdayStr)) // important line
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

	values := entities.AlertVariables{
		Category: category,
		Hours:    hour,
		Minutes:  minute,
		Days:     day,
		Weekdays: weekdayStr,
	}

	switch method {
	case "email":
		switch frequency {
		case "hourly":
			h.Deps.SpendAmount.HourlyNotification(category)
		case "daily":
			err := h.Deps.ManageAlerts.WriteNotificationValues(&values)
			if err != nil {
				log.Fatal(err)
			}
			h.Deps.SpendAmount.DailyNotification(hour, minute, category)
		case "weekly":
			h.Deps.SpendAmount.WeeklyNotification(weekday, hour, minute, category)
		case "monthly":
			h.Deps.SpendAmount.MonthlyNotification(day, hour, minute, category)
		default:
			return errors.New("wrong or no frequency is selected")
		}
	case "cli":
		return errors.New("you can't spend above your budget limit")
	}
	return nil
}

func (h MyAlert) WriteNotificationValues(av *entities.AlertVariables) error {
	hours := av.Hours
	minutes := av.Minutes

	desiredDay := av.Days
	desiredWeekday := av.Weekdays
	desiredTime := fmt.Sprintf("%02d:%02d", hours, minutes)

	// Convert to upper letter
	if len(desiredWeekday) > 0 {
		desiredWeekday = string(unicode.ToUpper(rune(desiredWeekday[0]))) + strings.ToLower(desiredWeekday[1:])
	}

	scriptPath := os.Getenv("SCRIPT_PATH")
	if scriptPath == "" {
		return errors.New("SCRIPT_PATH environment variable is not set")
	}
	scriptFile := fmt.Sprintf("%s/script.sh", scriptPath)

	input, err := os.ReadFile(scriptFile)
	if err != nil {
		log.Fatal(err)
	}

	value, err := h.Deps.AlertDB.ViewAlert(av.Category)
	if err != nil {
		return err
	}

	category, ok := value[1].(string)
	if !ok {
		return errors.New("unable to convert to string")
	}

	if len(av.Category) != 0 && av.Category == category {
		lines := strings.Split(string(input), "\n")
		for i, line := range lines {
			if strings.Contains(line, "DESIRED_TIME=") {
				lines[i] = fmt.Sprintf(`DESIRED_TIME="%s"`, desiredTime)
			}
			if strings.Contains(line, "DESIRED_DAY=") {
				lines[i] = fmt.Sprintf(`DESIRED_DAY="%d"`, desiredDay)
			}
			if strings.Contains(line, "DESIRED_WEEKDAY=") {
				lines[i] = fmt.Sprintf(`DESIRED_WEEKDAY="%s"`, desiredWeekday)
			}
		}

		output := strings.Join(lines, "\n")
		err = os.WriteFile(scriptFile, []byte(output), 0644)
		if err != nil {
			return err
		}
	} else {
		return errors.New("category is not present in the alert")
	}

	return nil
}

func (h MyAlert) SendNotification(category string) error {
	budgetDetails, err := h.Deps.ManageBudget.ViewBudget(category)
	if err != nil {
		return err
	}

	alertDetails, err := h.Deps.AlertDB.ViewAlert(category)
	if err != nil {
		return err
	}

	budgetCategory, ok1 := budgetDetails[1].(string)
	budgetAmount, ok2 := budgetDetails[2].(int)
	spentAmount, ok3 := budgetDetails[3].(int)
	alertCategory, ok4 := alertDetails[1].(string)

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return errors.New("unable to convert to int or string")
	}

	if len(category) != 0 {
		if category == budgetCategory && category == alertCategory {
			if spentAmount > budgetAmount {
				h.Deps.ManageAlerts.SendAlert(category)
			} else {
				fmt.Printf("The '%s' category amount is not exceeded\n", category)
			}
		} else {
			log.Fatal("Notification can't be sent. Either the category is not stored in budget or in alert")
		}
	} else {
		log.Fatal("Enter the category. See 'flow budget alert msg -h'")
	}
	return nil
}
