package alert_db

import (
	"github.com/ibilalkayy/flow/framework_drivers/db"
	"github.com/ibilalkayy/flow/interface_adapters"
)

type MyAlertDatabase struct {
	interface_adapters.AlertDatabase
	db.MyConnect
}
