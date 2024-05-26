package init

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/cmd"
	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/framework/db"
	"github.com/ibilalkayy/flow/handler"
	"github.com/ibilalkayy/flow/interfaces"
	usecases_init "github.com/ibilalkayy/flow/usecases/app/init"
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
	appPassword, _ := cmd.Flags().GetString("app-password")

	dbHost, _ := cmd.Flags().GetString("db-host")
	postgresPort, _ := cmd.Flags().GetString("db-port")
	postgresUser, _ := cmd.Flags().GetString("db-user")
	postgresPassword, _ := cmd.Flags().GetString("db-password")
	postgresDBName, _ := cmd.Flags().GetString("db-name")
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
		err := InitializeApplication(authParams, dbParams)
		if err != nil {
			fmt.Println("Error during initialization:", err)
			return
		}
	} else {
		fmt.Println(errors.New("please provide all the required flags"))
	}
}

func InitializeApplication(authParams *entities.AuthVariables, dbParams *entities.DatabaseVariables) error {
	myConnection := &db.MyConnection{}
	myInit := &usecases_init.MyInit{}
	deps := interfaces.Dependencies{
		Connect: myConnection,
		Init:    myInit,
	}
	handle := handler.NewHandler(deps)
	myInit.Handler = handle
	myConnection.Handler = handle

	err := handle.Deps.Init.WriteEnvFile(authParams, dbParams)
	if err != nil {
		return fmt.Errorf("error writing to .env file: %v", err)
	}

	_, err = handle.Deps.Connect.Connection()
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
	InitCmd.Flags().StringP("app-password", "a", "", "Write the App Password of your Gmail account. For more info visit: https://support.google.com/accounts/answer/185833")
	InitCmd.Flags().StringP("db-host", "o", "", "Write the PostgreSQL host")
	InitCmd.Flags().StringP("db-port", "p", "", "Write the PostgreSQL port")
	InitCmd.Flags().StringP("db-password", "w", "", "Write the PostgreSQL password")
	InitCmd.Flags().StringP("db-user", "u", "", "Write the PostgreSQL user")
	InitCmd.Flags().StringP("db-name", "d", "", "Write the PostgreSQL DB name")
	InitCmd.Flags().StringP("sslmode", "s", "", "Write the PostgreSQL SSLMode")
}
