package spend_subhandler

import (
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
)

func TakeHandler() *handler.Handler {
	myConnection := &db.MyConnection{}
	myHistory := &budget_db.MyBudgetDB{}

	deps := interfaces.Dependencies{
		Connect:      myConnection,
		ManageBudget: myHistory,
	}

	handle := handler.NewHandler(deps)
	myConnection.Handler = handle
	myHistory.Handler = handle

	return handle
}
