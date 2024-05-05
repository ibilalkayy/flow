package init

import (
	"errors"
	"fmt"

	"github.com/ibilalkayy/flow/cmd"
	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/framework_drivers/db"
	"github.com/spf13/cobra"
)

func (MyInit) AllNonEmpty(params ...string) bool {
	for _, param := range params {
		if len(param) == 0 {
			return false
		}
	}
	return true
}

func Initialize(cmd *cobra.Command, args []string) {
	var m MyInit
	m.InitApp(cmd, args)
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your flow application",
	Run:   Initialize,
}

func (m MyInit) InitApp(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("username")
	gmail, _ := cmd.Flags().GetString("gmail")
	appPassword, _ := cmd.Flags().GetString("app_password")

	dbHost, _ := cmd.Flags().GetString("host")
	postgresPort, _ := cmd.Flags().GetString("port")
	postgresUser, _ := cmd.Flags().GetString("db_username")
	postgresPassword, _ := cmd.Flags().GetString("db_password")
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

	if m.AllNonEmpty(
		authParams.Username, authParams.Gmail, authParams.AppPassword,
		dbParams.Host, dbParams.Port, dbParams.User, dbParams.Password,
		dbParams.DBName, dbParams.SSLMode,
	) {
		err := m.InitializeApplication(authParams, dbParams)
		if err != nil {
			fmt.Println("Error during initialization:", err)
			return
		}
	} else {
		fmt.Println(errors.New("please provide all the required flags"))
	}
}

func (m MyInit) InitializeApplication(authParams *entities.AuthVariables, dbParams *entities.DatabaseVariables) error {
	err := m.WriteEnvFile(authParams, dbParams)
	if err != nil {
		return fmt.Errorf("error writing to .env file: %v", err)
	}

	var c db.MyConnect
	_, err = c.Connection()
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
	InitCmd.Flags().StringP("db_password", "w", "", "Write the PostgreSQL password")
	InitCmd.Flags().StringP("db_username", "u", "", "Write the PostgreSQL username")
	InitCmd.Flags().StringP("dbname", "d", "", "Write the PostgreSQL DB name")
	InitCmd.Flags().StringP("sslmode", "s", "", "Write the PostgreSQL SSLMode")
}
