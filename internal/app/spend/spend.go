package internal_spending

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/db/budget_db"
	"github.com/ibilalkayy/flow/email"
)

func SpendMoney(category string, spending_amount int) error {
	var answer string
	values, err := budget_db.ViewBudget(category)
	if err != nil {
		return err
	}

	categoryName, ok1 := values[1].(string)
	remainingAmount, ok2 := values[4].(int)

	if !ok1 || !ok2 {
		return errors.New("unable to convert budget amount to int or string")
	}

	if category == categoryName {
		if spending_amount <= remainingAmount {
			err := budget_db.AddExpenditure(spending_amount, category)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("You have spent %d more than your set budget\n", remainingAmount)
			fmt.Printf("Do you still want to spend? [yes/no]: ")
			fmt.Scanln(&answer)

			switch answer {
			case "yes", "y":
				email.SendAlertEmail(category)
				err := budget_db.AddExpenditure(spending_amount, category)
				if err != nil {
					return err
				}
				fmt.Println("Enjoy your spending!")
			case "no", "n":
				fmt.Println("Alright")
			default:
				return errors.New("select the right option")
			}
		}
	} else {
		return errors.New("category is not found. first setup the alert. see 'flow budget alert setup -h'")
	}
	return nil
}
