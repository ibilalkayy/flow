package main

import (
	"github.com/ibilalkayy/flow/cmd"
	_ "github.com/ibilalkayy/flow/cmd/budget"
	_ "github.com/ibilalkayy/flow/cmd/init"
	_ "github.com/ibilalkayy/flow/cmd/transaction"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cmd.Execute()
}
