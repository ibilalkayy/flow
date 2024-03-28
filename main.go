package main

import (
	"github.com/ibilalkayy/flow/cmd"
	_ "github.com/ibilalkayy/flow/cmd/budget"
	_ "github.com/ibilalkayy/flow/cmd/init"
	_ "github.com/ibilalkayy/flow/cmd/spend"
)

func main() {
	cmd.Execute()
}
