package middleware

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Println("add credentials to the .env file")
	}

	return value
}
