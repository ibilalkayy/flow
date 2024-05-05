package total_amount_db

import (
	"github.com/ibilalkayy/flow/framework_drivers/db"
	"github.com/ibilalkayy/flow/interface_adapters"
)

type MyTotalDatabase struct {
	interface_adapters.TotalAmountDatabase
	db.MyConnect
}
