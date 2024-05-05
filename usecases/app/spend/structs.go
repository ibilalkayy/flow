package usecases_spending

import (
	"github.com/ibilalkayy/flow/framework_drivers/db/budget_db"
	"github.com/ibilalkayy/flow/framework_drivers/db/total_amount_db"
	"github.com/ibilalkayy/flow/framework_drivers/email"
	"github.com/ibilalkayy/flow/interface_adapters"
)

type MySpending struct {
	email.MyEmail
	total_amount_db.MyTotalDatabase
	budget_db.MyBudgetDatabase
	interface_adapters.Spending
	interface_adapters.Notifications
	interface_adapters.History
	budget_db.MyHistoryDatabase
}
