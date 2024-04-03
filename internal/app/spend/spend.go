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

	budgetAmount, ok := values[2].(int)
	if !ok {
		return errors.New("unable to convert budget amount to int")
	}

	if category == values[1] {
		if spending_amount <= budgetAmount {
			budget_db.UpdateBudget(category, "", 0, spending_amount, budgetAmount-spending_amount)
			fmt.Println("Enjoy your spending!")
		} else {
			fmt.Printf("Your spending amount is exceeded. Do you still want to continue? [yes/no]: ")
			fmt.Scanln(&answer)

			switch answer {
			case "yes", "y":
				email.SendAlertEmail(category)
				budget_db.UpdateBudget(category, "", 0, spending_amount, budgetAmount-spending_amount)
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
