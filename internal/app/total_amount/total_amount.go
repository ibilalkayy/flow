package internal_total_amount

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/db"
	"github.com/ibilalkayy/flow/db/total_amount_db"
	"github.com/ibilalkayy/flow/internal/common/structs"
)

func SetTotalAmount(totalAmount int, include_category, label string) error {
	tav := structs.TotalAmountVariables{
		TotalAmount:     totalAmount,
		RemainingAmount: 0,
		Status:          "inactive",
	}

	tacv := structs.TotalAmountVariables{
		Included: include_category,
		Label:    label,
	}

	amountExists, err := db.TableExists("totalamount")
	if err != nil {
		return err
	}

	categoryExists, err := db.TableExists("totalamountcategory")
	if err != nil {
		return err
	}

	if amountExists && categoryExists {
		err := handleExistingTables(totalAmount, tav, tacv)
		if err != nil {
			return err
		}
	} else {
		err := handleMissingTables(tav, tacv)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleExistingTables(totalAmount int, tav, tacv structs.TotalAmountVariables) error {
	_, values, err := total_amount_db.ViewTotalAmountCategory()
	if err != nil {
		return err
	}

	amount, err := total_amount_db.ViewTotalAmount()
	if err != nil {
		return err
	}

	total_amount, ok := amount[1].(int)
	if !ok {
		return errors.New("unable to convert string to int")
	}

	if len(values) == 0 {
		err = total_amount_db.InsertTotalAmount(&tav, "db/migrations/")
		if err != nil {
			return err
		}

		err = total_amount_db.InsertTotalAmountCategory(&tacv, "db/migrations/")
		if err != nil {
			return err
		}
	} else {
		if total_amount != 0 && totalAmount != 0 {
			return errors.New("total amount is already set, now only add categories and labels")
		} else {
			for _, list := range values {
				if len(list[0]) != 0 && len(list[1]) != 0 {
					err = total_amount_db.InsertTotalAmountCategory(&tacv, "db/migrations/")
					if err != nil {
						return err
					}
					fmt.Println("Category and label is successfully included!")
					break
				}
			}
		}
	}
	return nil
}

func handleMissingTables(tav, tacv structs.TotalAmountVariables) error {
	err := total_amount_db.InsertTotalAmount(&tav, "db/migrations/")
	if err != nil {
		return err
	}

	err = total_amount_db.InsertTotalAmountCategory(&tacv, "db/migrations/")
	if err != nil {
		return err
	}
	return nil
}
