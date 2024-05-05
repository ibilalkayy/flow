package usecases_total_amount

import (
	"github.com/ibilalkayy/flow/framework_drivers/db"
	"github.com/ibilalkayy/flow/framework_drivers/db/total_amount_db"
)

type MyTotalAmount struct {
	total_amount_db.MyTotalDatabase
	db.MyConnect
}
