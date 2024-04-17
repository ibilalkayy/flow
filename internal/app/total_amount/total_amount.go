package internal_total_amount

import (
	"errors"

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

		for _, list := range values {
			if len(list[0]) != 0 && len(list[1]) != 0 && total_amount != 0 {
				return errors.New("you've already added the total amount. now add only categories and labels without writing the total amount")
			}
		}
	} else {
		err = total_amount_db.InsertTotalAmount(&tav, "db/migrations/")
		if err != nil {
			return err
		}

		err = total_amount_db.InsertTotalAmountCategory(&tacv, "db/migrations/")
		if err != nil {
			return err
		}
	}
	return nil
}
