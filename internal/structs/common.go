package structs

import (
	"log"
	"strconv"
)

type AuthVariables struct {
	Username    string
	Gmail       string
	AppPassword string
}

type DatabaseVariables struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type BudgetVariables struct {
	Category string
	Amount   int
	Spent    int
}

type AlertVariables struct {
	Category  string
	Frequency string
	Method    string
	Days      int
	Weekdays  string
	Hours     int
	Minutes   int
	Seconds   int
}

type EmailVariables struct {
	Username       string
	Category       string
	CategoryAmount int
}

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
