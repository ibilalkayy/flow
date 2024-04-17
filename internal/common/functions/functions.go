package functions

import (
	"errors"
	"log"
	"strconv"

	"github.com/ibilalkayy/flow/db/total_amount_db"
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

func TotalAmountValues() ([][2]string, int, string, error) {
	values, err := total_amount_db.ViewTotalAmount()
	if err != nil {
		return [][2]string{}, 0, "", err
	}

	_, includedCategory, err := total_amount_db.ViewTotalAmountCategory()
	if err != nil {
		return [][2]string{}, 0, "", err
	}

	totalAmount, ok1 := values[1].(int)
	status, ok2 := values[2].(string)

	if !ok1 || !ok2 {
		return [][2]string{}, 0, "", errors.New("unable to convert to int or string")
	}

	return includedCategory, totalAmount, status, nil
}
