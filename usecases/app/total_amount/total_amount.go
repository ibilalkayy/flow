package usecases_total_amount

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/handler"
)

type MyTotalAmount struct {
	*handler.Handler
}

func (h MyTotalAmount) SetTotalAmount(totalAmount int, include_category, label string) error {
	tav := entities.TotalAmountVariables{
		TotalAmount:     totalAmount,
		SpentAmount:     0,
		RemainingAmount: 0,
		Status:          "inactive",
	}

	tacv := entities.TotalAmountVariables{
		Included: include_category,
		Label:    label,
	}

	amountExists, err := h.Deps.Connect.TableExists("Totalamount")
	if err != nil {
		return err
	}

	categoryExists, err := h.Deps.Connect.TableExists("Totalamountcategories")
	if err != nil {
		return err
	}

	if amountExists && categoryExists {
		err := h.HandleExistingTables(totalAmount, tav, tacv)
		if err != nil {
			return err
		}
	} else {
		err := h.HandleMissingTables(tav, tacv)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h MyTotalAmount) HandleExistingTables(totalAmount int, tav, tacv entities.TotalAmountVariables) error {
	_, values, err := h.Deps.TotalAmountCategory.ViewTotalAmountCategories()
	if err != nil {
		return err
	}

	amount, err := h.Deps.TotalAmount.ViewTotalAmount()
	if err != nil {
		return err
	}

	total_amount, ok := amount[1].(int)
	if !ok {
		return errors.New("unable to convert string to int")
	}

	if len(values) == 0 {
		err = h.Deps.TotalAmount.InsertTotalAmount(&tav)
		if err != nil {
			return err
		}

		err = h.Deps.TotalAmountCategory.InsertTotalAmountCategory(&tacv)
		if err != nil {
			return err
		}
	} else {
		if total_amount != 0 && totalAmount != 0 {
			return errors.New("total amount is already set, now only add categories and labels")
		} else {
			for _, list := range values {
				if len(list[0]) != 0 && len(list[1]) != 0 {
					err = h.Deps.TotalAmountCategory.InsertTotalAmountCategory(&tacv)
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

func (h MyTotalAmount) HandleMissingTables(tav, tacv entities.TotalAmountVariables) error {
	err := h.Deps.TotalAmount.InsertTotalAmount(&tav)
	if err != nil {
		return err
	}

	err = h.Deps.TotalAmountCategory.InsertTotalAmountCategory(&tacv)
	if err != nil {
		return err
	}
	return nil
}
