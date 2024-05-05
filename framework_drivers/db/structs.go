package db

import (
	"github.com/ibilalkayy/flow/interface_adapters"
	"github.com/ibilalkayy/flow/usecases/middleware"
)

type MyConnect struct {
	interface_adapters.Connect
	middleware.LoadEnv
}
