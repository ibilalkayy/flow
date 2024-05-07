package total_amount_db

import (
	conversion "github.com/ibilalkayy/flow/common"
	"github.com/ibilalkayy/flow/framework_drivers/db"
	"github.com/ibilalkayy/flow/interface_adapters"
)

type MyTotalDatabase struct {
	interface_adapters.BudgetDatabase
	interface_adapters.TotalAmountDatabase
	conversion.MyConversion
	db.MyConnect
}
