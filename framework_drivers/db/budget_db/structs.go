package budget_db

import (
	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/framework_drivers/db"
	"github.com/ibilalkayy/flow/framework_drivers/db/total_amount_db"
	"github.com/ibilalkayy/flow/interface_adapters"
)

type MyBudgetDatabase struct {
	interface_adapters.BudgetDatabase
	total_amount_db.MyTotalDatabase
	db.MyConnect
	conversion.MyConversion
}

type MyHistoryDatabase struct {
	interface_adapters.HistoryDatabase
	total_amount_db.MyTotalDatabase
	db.MyConnect
}
