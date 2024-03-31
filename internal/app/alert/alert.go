package internal_alert

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ibilalkayy/flow/db/alert_db"
	internal_budget "github.com/ibilalkayy/flow/internal/app/budget"
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
