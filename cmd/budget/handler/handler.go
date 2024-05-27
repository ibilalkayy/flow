package budget_handler

import (
	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
)

func TakeHandler() *handler.Handler {
	myConnection := &db.MyConnection{}
	myTotalDB := &total_amount_db.MyTotalAmountDB{}
	myBudgetDB := &budget_db.MyBudgetDB{}
	myCommon := &conversion.MyCommon{}

	deps := interfaces.Dependencies{
		Connect:             myConnection,
		TotalAmount:         myTotalDB,
		TotalAmountCategory: myTotalDB,
		ManageBudget:        myBudgetDB,
		Common:              myCommon,
	}

	handle := handler.NewHandler(deps)
	myConnection.Handler = handle
	myTotalDB.Handler = handle
	myBudgetDB.Handler = handle
	myCommon.Handler = handle

	return handle
}
