package internal_spending

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/db/budget_db"
	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/ibilalkayy/flow/email"
	"github.com/ibilalkayy/flow/internal/common/functions"
)

func SpendMoney(category string, spending_amount int) error {
	var answer string
	values, err := budget_db.ViewBudget(category)
	if err != nil {
		return err
	}

	categoryName, ok1 := values[1].(string)
	budget_category_amount, ok2 := values[2].(int)
	budget_category_spent_amount, ok3 := values[3].(int)
	budget_remaining_amount, ok4 := values[4].(int)
	included_categories_in_total_amount, total_amount, total_spent_amount, total_amount_status, err := functions.TotalAmountValues()
	if err != nil {
		return err
	}

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return errors.New("unable to convert budget amount to int or string")
	}

	if total_amount_status != "active" {
		return errors.New("make your total amount status active. see 'flow total-amount -h'")
	}

	if total_spent_amount+spending_amount > total_amount {
		return errors.New("you have exceeded the total amount")
	}

	foundCategory := false
	for _, list := range included_categories_in_total_amount {
		if category == categoryName && category == list[0] {

			foundCategory = true
			budget_category_total_spending_amount := spending_amount + budget_category_spent_amount

			if budget_category_total_spending_amount <= budget_category_amount {
				err := budget_db.AddBudgetExpenditure(spending_amount, category)
				if err != nil {
					return err
				}
				err = total_amount_db.CalculateRemaining(category)
				if err != nil {
					return err
				}
				fmt.Println("Enjoy your spending!")
				break
			} else if spending_amount <= budget_remaining_amount {
				err := budget_db.AddBudgetExpenditure(spending_amount, category)
				if err != nil {
					return err
				}
				err = total_amount_db.CalculateRemaining(category)
				if err != nil {
					return err
				}
				fmt.Println("Enjoy your spending!")
				break
			} else if spending_amount > budget_remaining_amount && spending_amount <= total_amount && budget_category_spent_amount <= total_amount && budget_category_total_spending_amount <= total_amount {
				fmt.Printf("You have spent %d and your remaining balance is %d but your budget is %d\n", budget_category_spent_amount, budget_remaining_amount, budget_category_amount)
				fmt.Printf("Do you still want to spend? [yes/no]: ")
				fmt.Scanln(&answer)

				switch answer {
				case "yes", "y":
					email.SendAlertEmail(category)
					err := budget_db.AddBudgetExpenditure(spending_amount, category)
					if err != nil {
						return err
					}
					err = total_amount_db.CalculateRemaining(category)
					if err != nil {
						return err
					}
					fmt.Println("Enjoy your spending!")
				case "no", "n":
					fmt.Println("Alright")
				default:
					return errors.New("select the right option")
				}
				break
			} else {
				return errors.New("you have exceeded the total amount")
			}
		}
	}

	if !foundCategory {
		return errors.New("category is not found. setup the alert 'flow budget alert setup -h' or include the category in your total amount 'flow total-amount set -h'")
	}

	return nil
}
