package conversion

import (
	"errors"
	"log"
	"strconv"

	"github.com/ibilalkayy/flow/framework_drivers/db/total_amount_db"
)

func IntToString(key int) string {
	value := strconv.Itoa(key)
	return value
}

func StringToInt(key string) int {
	if key == "" {
		return 0
	}
	value, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func TotalAmountValues() ([][2]string, [3]interface{}, error) {
	values, err := total_amount_db.ViewTotalAmount()
	if err != nil {
		return [][2]string{}, [3]interface{}{}, err
	}

	_, includedCategory, err := total_amount_db.ViewTotalAmountCategory()
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
