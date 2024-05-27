package budget_subhandler

import (
	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/alert_db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/framework/email"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	usecases_alert "github.com/ibilalkayy/flow/usecases/app/alert"
	usecases_budget "github.com/ibilalkayy/flow/usecases/app/budget"
	usecases_spending "github.com/ibilalkayy/flow/usecases/app/spend"
	"github.com/ibilalkayy/flow/usecases/middleware"
)

func TakeHandler() *handler.Handler {
	myConnection := &db.MyConnection{}
	myAlert := &usecases_alert.MyAlert{}
	myAlertDB := &alert_db.MyAlertDB{}
	myBudgetDB := &budget_db.MyBudgetDB{}
	myBudget := &usecases_budget.MyBudget{}
	mySpending := &usecases_spending.MySpending{}
	myCommon := &conversion.MyCommon{}
	myEmail := &email.MyEmail{}
	myEnv := &middleware.MyEnv{}

	deps := interfaces.Dependencies{
		Connect:      myConnection,
		AlertDB:      myAlertDB,
		ManageAlerts: myAlert,
		Budget:       myBudget,
		ManageBudget: myBudgetDB,
		SpendAmount:  mySpending,
		Common:       myCommon,
		HandleEmail:  myEmail,
		Env:          myEnv,
	}

	handle := handler.NewHandler(deps)
	myConnection.Handler = handle
	myAlertDB.Handler = handle
	myAlert.Handler = handle
	myBudgetDB.Handler = handle
	myBudget.Handler = handle
	mySpending.Handler = handle
	myEmail.Handler = handle
	myCommon.Handler = handle
	myEnv.Handler = handle

	return handle
}
