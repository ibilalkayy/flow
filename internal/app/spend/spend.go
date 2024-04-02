package internal_spending

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/db/budget_db"
	"github.com/ibilalkayy/flow/email"
)

func SpendMoney(category, spending_amount string) error {
	var answer string
	values, err := budget_db.ViewBudget(category)
	if err != nil {
		return err
	}

	if category == values[1] {
		if spending_amount <= values[2] {
			budget_db.UpdateBudget(category, "", "", spending_amount)
			fmt.Println("Enjoy your spending!")
		} else {
			fmt.Printf("Your spending amount is exceeded. Do you still want to continue? [yes/no]: ")
			fmt.Scanln(&answer)

			switch answer {
			case "yes", "y":
				budget_db.UpdateBudget(category, "", "", spending_amount)
				email.SendAlertEmail(category)
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
