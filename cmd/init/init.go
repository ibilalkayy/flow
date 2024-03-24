package init

import (
	"fmt"

	"github.com/ibilalkayy/flow/cmd"
	"github.com/ibilalkayy/flow/db/budget_db"
	internal_init "github.com/ibilalkayy/flow/internal/app/init"
	"github.com/ibilalkayy/flow/internal/structs"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your flow application",
	Run: func(cmd *cobra.Command, args []string) {
		db_host, _ := cmd.Flags().GetString("host")
		postgres_port, _ := cmd.Flags().GetString("port")
		postgres_user, _ := cmd.Flags().GetString("user")
		postgres_password, _ := cmd.Flags().GetString("password")
		postgres_dbname, _ := cmd.Flags().GetString("dbname")
		sslmode, _ := cmd.Flags().GetString("sslmode")

		parameters := &structs.DatabaseVariables{
			Host:     db_host,
			Port:     postgres_port,
			User:     postgres_user,
			Password: postgres_password,
			DBName:   postgres_dbname,
			SSLMode:  sslmode,
		}

		err := internal_init.WriteEnvFile(parameters)
		if err != nil {
			fmt.Println("Error writing to .env file:", err)
			return
		}

		_, err = budget_db.Connection()
		if err != nil {
			fmt.Println("Error connecting to the database:", err)
			return
		}

		fmt.Println("Successfully connected to the database and saved values to .env file!")
	},
}

func init() {
	cmd.RootCmd.AddCommand(InitCmd)
	InitCmd.Flags().StringP("host", "o", "", "Write the PostgreSQL host")
	InitCmd.Flags().StringP("port", "p", "", "Write the PostgreSQL port")
	InitCmd.Flags().StringP("password", "w", "", "Write the PostgreSQL password")
	InitCmd.Flags().StringP("user", "u", "", "Write the PostgreSQL username")
	InitCmd.Flags().StringP("dbname", "d", "", "Write the PostgreSQL DB name")
	InitCmd.Flags().StringP("sslmode", "s", "", "Write the PostgreSQL SSLMode")
}
