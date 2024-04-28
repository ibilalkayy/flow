package init

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/cmd"
	"github.com/ibilalkayy/flow/db"
	internal_init "github.com/ibilalkayy/flow/internal/app/init"
	"github.com/ibilalkayy/flow/internal/entities"
	"github.com/spf13/cobra"
)

func allNonEmpty(params ...string) bool {
	for _, param := range params {
		if len(param) == 0 {
			return false
		}
	}
	return true
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your flow application",
	Run:   initApp,
}

func initApp(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("username")
	gmail, _ := cmd.Flags().GetString("gmail")
	appPassword, _ := cmd.Flags().GetString("app_password")

	dbHost, _ := cmd.Flags().GetString("host")
	postgresPort, _ := cmd.Flags().GetString("port")
	postgresUser, _ := cmd.Flags().GetString("user")
	postgresPassword, _ := cmd.Flags().GetString("password")
	postgresDBName, _ := cmd.Flags().GetString("dbname")
	sslMode, _ := cmd.Flags().GetString("sslmode")

	authParams := &entities.AuthVariables{
		Username:    username,
		Gmail:       gmail,
		AppPassword: appPassword,
	}

	dbParams := &entities.DatabaseVariables{
		Host:     dbHost,
		Port:     postgresPort,
		User:     postgresUser,
		Password: postgresPassword,
		DBName:   postgresDBName,
		SSLMode:  sslMode,
	}

	if allNonEmpty(
		authParams.Username, authParams.Gmail, authParams.AppPassword,
		dbParams.Host, dbParams.Port, dbParams.User, dbParams.Password,
		dbParams.DBName, dbParams.SSLMode,
	) {
		err := initializeApplication(authParams, dbParams)
		if err != nil {
			fmt.Println("Error during initialization:", err)
			return
		}
	} else {
		fmt.Println(errors.New("please provide all the required flags"))
	}
}

func initializeApplication(authParams *entities.AuthVariables, dbParams *entities.DatabaseVariables) error {
	err := internal_init.WriteEnvFile(authParams, dbParams)
	if err != nil {
		return fmt.Errorf("error writing to .env file: %v", err)
	}

	_, err = db.Connection()
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")
	return nil
}

func init() {
	cmd.RootCmd.AddCommand(InitCmd)
	InitCmd.Flags().StringP("username", "n", "", "Write your username")
	InitCmd.Flags().StringP("gmail", "g", "", "Write your Gmail address for alert notifications")
	InitCmd.Flags().StringP("app_password", "a", "", "Write the App Password of your Gmail account. For more info visit: https://support.google.com/accounts/answer/185833")
	InitCmd.Flags().StringP("host", "o", "", "Write the PostgreSQL host")
	InitCmd.Flags().StringP("port", "p", "", "Write the PostgreSQL port")
	InitCmd.Flags().StringP("password", "w", "", "Write the PostgreSQL password")
	InitCmd.Flags().StringP("user", "u", "", "Write the PostgreSQL username")
	InitCmd.Flags().StringP("dbname", "d", "", "Write the PostgreSQL DB name")
	InitCmd.Flags().StringP("sslmode", "s", "", "Write the PostgreSQL SSLMode")
}
