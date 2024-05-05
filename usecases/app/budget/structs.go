package usecases_budget

import (
	"github.com/ibilalkayy/flow/framework_drivers/db"
	"github.com/ibilalkayy/flow/interface_adapters"
)

type MyBudget struct {
	interface_adapters.Budget
	db.MyConnect
}
