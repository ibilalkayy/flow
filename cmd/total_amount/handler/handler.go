package total_amount_handler

import (
	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/budget_db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	usecases_total_amount "github.com/ibilalkayy/flow/usecases/app/total_amount"
)

func TakeHandler() *handler.Handler {
	myConnection := &db.MyConnection{}
	myTotalDB := &total_amount_db.MyTotalAmountDB{}
	myTotal := &usecases_total_amount.MyTotalAmount{}
	myBudget := &budget_db.MyBudgetDB{}
	myCommon := &conversion.MyCommon{}

	deps := interfaces.Dependencies{
		Connect:             myConnection,
		TotalAmount:         myTotalDB,
		TotalAmountCategory: myTotalDB,
		Total:               myTotal,
		ManageBudget:        myBudget,
		Common:              myCommon,
	}

	handle := handler.NewHandler(deps)
	myConnection.Handler = handle
	myTotal.Handler = handle
	myTotalDB.Handler = handle
	myBudget.Handler = handle
	myCommon.Handler = handle

	return handle
}
