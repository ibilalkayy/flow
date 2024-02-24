package main

import (
	"github.com/ibilalkayy/flow/cmd"
	_ "github.com/ibilalkayy/flow/cmd/budget"
	_ "github.com/ibilalkayy/flow/cmd/budget/adjust"
	_ "github.com/ibilalkayy/flow/cmd/budget/create"
	_ "github.com/ibilalkayy/flow/cmd/budget/remove"
	_ "github.com/ibilalkayy/flow/cmd/budget/view"
)

func main() {
	cmd.Execute()
}
