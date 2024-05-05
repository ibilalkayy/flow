package total_amount_db

import "errors"

func (m MyTotalDatabase) TotalAmountValues() ([][2]string, [3]interface{}, error) {
	values, err := m.ViewTotalAmount()
	if err != nil {
		return [][2]string{}, [3]interface{}{}, err
	}

	_, includedCategory, err := m.ViewTotalAmountCategory()
	if err != nil {
		return [][2]string{}, [3]interface{}{}, err
	}

	totalAmount, ok1 := values[1].(int)
	spentAmount, ok2 := values[2].(int)
	status, ok3 := values[4].(string)

	if !ok1 || !ok2 || !ok3 {
		return [][2]string{}, [3]interface{}{}, errors.New("unable to convert to int or string")
	}

	details := [3]interface{}{totalAmount, spentAmount, status}
	return includedCategory, details, nil
}
