package init

import (
	"github.com/ibilalkayy/flow/interface_adapters"
	usecases_init "github.com/ibilalkayy/flow/usecases/app/init"
)

type MyInit struct {
	interface_adapters.Init
	usecases_init.MyEnvFile
}
