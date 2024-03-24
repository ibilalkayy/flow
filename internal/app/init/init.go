package internal_init

import (
	"fmt"
	"os"

	"github.com/ibilalkayy/flow/internal/structs"
)

func WriteEnvFile(dv *structs.DatabaseVariables) error {
	f, err := os.Create(".env")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "DB_HOST=%s\n", dv.Host)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "DB_PORT=%s\n", dv.Port)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "DB_USER=%s\n", dv.User)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "DB_PASSWORD=%s\n", dv.Password)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "DB_NAME=%s\n", dv.DBName)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "SSL_MODE=%s\n", dv.SSLMode)
	if err != nil {
		return err
	}

	return nil
}
