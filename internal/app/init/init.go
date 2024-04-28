package internal_init

import (
	"fmt"
	"os"

	"github.com/ibilalkayy/flow/internal/entities"
)

func WriteEnvFile(av *entities.AuthVariables, dv *entities.DatabaseVariables) error {
	f, err := os.Create(".env")
	if err != nil {
		return err
	}
	defer f.Close()

	fields := []struct {
		Key   string
		Value string
	}{
		{"USERNAME", av.Username},
		{"APP_EMAIL", av.Gmail},
		{"APP_PASSWORD", av.AppPassword},

		{"DB_HOST", dv.Host},
		{"DB_PORT", dv.Port},
		{"DB_USER", dv.User},
		{"DB_PASSWORD", dv.Password},
		{"DB_NAME", dv.DBName},
		{"SSL_MODE", dv.SSLMode},
	}

	for _, field := range fields {
		if _, err := fmt.Fprintf(f, "%s=%s\n", field.Key, field.Value); err != nil {
			return err
		}
	}

	return nil
}
