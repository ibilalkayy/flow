package conversion

import (
	"log"
	"strconv"

	"github.com/ibilalkayy/flow/handler"
)

type MyCommon struct {
	*handler.Handler
}

func (MyCommon) IntToString(key int) string {
	value := strconv.Itoa(key)
	return value
}

func (MyCommon) StringToInt(key string) int {
	if key == "" {
		return 0
	}
	value, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
