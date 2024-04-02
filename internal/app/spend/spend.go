package internal_spending

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibilalkayy/flow/db/alert_db"
	"github.com/ibilalkayy/flow/db/spend_db"
	"github.com/ibilalkayy/flow/email"
	"github.com/ibilalkayy/flow/internal/structs"
)

func InsertSpending(value [3]string, exceeded string) {
	data := structs.SpendingVariables{
		Category:       value[0],
		CategoryAmount: value[1],
		SpendingAmount: value[2],
	}

	err := spend_db.CreateSpending(&data, exceeded, "db/migrations/")
	if err != nil {
		log.Fatal(err)
	}
}

func SpendMoney(category, spending_amount string) error {
	var answer string
	values, err := alert_db.ViewAlert(category)
	if err != nil {
		return err
	}

	value := [3]string{values[0], values[1], spending_amount}
	if category == values[0] {
		if spending_amount <= values[1] {
			InsertSpending(value, "false")
			fmt.Println("Enjoy your spending!")
		} else {
			fmt.Printf("Your spending amount is exceeded. Do you still want to continue? [yes/no]: ")
			fmt.Scanln(&answer)

			switch answer {
			case "yes", "y":
				InsertSpending(value, "true")
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
