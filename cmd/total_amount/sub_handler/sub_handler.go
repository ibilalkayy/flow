package total_amount_subhandler

import (
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/framework/db/total_amount_db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
)

func TakeHandler() *handler.Handler {
	myConnection := &db.MyConnection{}
	myTotalDB := &total_amount_db.MyTotalAmountDB{}
	deps := interfaces.Dependencies{
		Connect:             myConnection,
		TotalAmount:         myTotalDB,
		TotalAmountCategory: myTotalDB,
	}

	handle := handler.NewHandler(deps)
	myConnection.Handler = handle
	myTotalDB.Handler = handle

	return handle
}
