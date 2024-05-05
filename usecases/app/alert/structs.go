package usecases_alert

import (
	"github.com/ibilalkayy/flow/framework_drivers/db/alert_db"
	"github.com/ibilalkayy/flow/framework_drivers/db/budget_db"
	"github.com/ibilalkayy/flow/interface_adapters"
	usecases_budget "github.com/ibilalkayy/flow/usecases/app/budget"
	usecases_spending "github.com/ibilalkayy/flow/usecases/app/spend"
)

type MyAlerts struct {
	interface_adapters.Alerts
	usecases_budget.MyBudget
	alert_db.MyAlertDatabase
	budget_db.MyBudgetDatabase
	usecases_spending.MySpending
}
