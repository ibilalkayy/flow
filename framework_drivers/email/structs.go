package email

import (
	"github.com/ibilalkayy/flow/framework_drivers/db/budget_db"
	"github.com/ibilalkayy/flow/interface_adapters"
	"github.com/ibilalkayy/flow/usecases/middleware"
)

type MyEmail struct {
	interface_adapters.Email
	middleware.LoadEnv
	budget_db.MyBudgetDatabase
}
