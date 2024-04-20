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
	totalAmount, ok2 := values[2].(int)
	spentAmount, ok3 := values[3].(int)
	remainingAmount, ok4 := values[4].(int)
	totalAmountIncludedCategory, totalAllocatedAmount, totalAmountStatus, err := functions.TotalAmountValues()
	if err != nil {
		return err
	}

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return errors.New("unable to convert budget amount to int or string")
	}

	if totalAmountStatus != "active" {
		return errors.New("make your total amount status active. see 'flow total-amount -h'")
	}

	foundCategory := false
	for _, list := range totalAmountIncludedCategory {
		if category == categoryName && category == list[0] {
			foundCategory = true
			fmt.Println(list[0])

			totalSpent := spending_amount + spentAmount
			if totalSpent <= totalAmount {
				err := budget_db.AddExpenditure(spending_amount, category)
				if err != nil {
					return err
				}
				err = total_amount_db.CalculateRemaining(category)
				if err != nil {
					return err
				}
				fmt.Println("Enjoy your spending!")
			} else if spending_amount <= remainingAmount {
				err := budget_db.AddExpenditure(spending_amount, category)
				if err != nil {
					return err
				}
				err = total_amount_db.CalculateRemaining(category)
				if err != nil {
					return err
				}
				fmt.Println("Enjoy your spending!")
			} else if spending_amount > remainingAmount && spending_amount <= totalAllocatedAmount && spentAmount <= totalAllocatedAmount && totalSpent <= totalAllocatedAmount {
				fmt.Printf("You have spent %d and your remaining balance is %d but your budget is %d\n", spentAmount, remainingAmount, totalAmount)
				fmt.Printf("Do you still want to spend? [yes/no]: ")
				fmt.Scanln(&answer)

				switch answer {
				case "yes", "y":
					email.SendAlertEmail(category)
					err := budget_db.AddExpenditure(spending_amount, category)
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
