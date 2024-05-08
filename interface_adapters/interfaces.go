package interface_adapters

import (
	"database/sql"
	"time"

	"github.com/ibilalkayy/flow/entities"
	"github.com/spf13/cobra"
)

type Init interface {
	AllNonEmpty(params ...string) bool
	InitApp(cmd *cobra.Command, args []string)
	InitializeApplication(authParams *entities.AuthVariables, dbParams *entities.DatabaseVariables) error
}

type Connect interface {
	Connection() (*sql.DB, error)
	Table(filename string, number int) (*sql.DB, error)
	TableExists(tableName string) (bool, error)
}

type TotalAmountDatabase interface {
	// Amount
	InsertTotalAmount(tv *entities.TotalAmountVariables) error
	ViewTotalAmount() ([5]interface{}, error)
	RemoveTotalAmount(category string) error
	UpdateTotalAmount(tv *entities.TotalAmountVariables) error
	UpdateStatus(tv *entities.TotalAmountVariables) error
	CalculateRemaining(category string) error

	// Categories
	InsertTotalAmountCategory(tv *entities.TotalAmountVariables) error
	ViewTotalAmountCategories() (string, [][2]string, error)

	// values
	TotalAmountValues() ([][2]string, [3]interface{}, error)
}

type BudgetDatabase interface {
	CreateBudget(bv *entities.BudgetVariables) error
	TakeBudgetAmounts() ([]int, error)
	ViewBudget(category string) ([5]interface{}, error)
	RemoveBudget(category string) error
	UpdateBudget(old, new string, amount int) error
	AddBudgtExpenditure(spent int, category string) error
	GetBudgetData(filepath, filename string) error
}

type Budget interface {
	CategoryAmount(category string) (string, int, error)
}

type HistoryDatabase interface {
	InsertHistory(hv *entities.HistoryVariables) error
	ViewHistory(category string) ([2]interface{}, error)
	RemoveHistory(category string) error
}

type History interface {
	StoreHistory(category string, spending_amount int) error
}

type Spending interface {
	SpendMoney(category string, spending_amount int) error
}

type Notifications interface {
	HourlyNotification(category string)
	DailyNotification(hour, min, sec int, category string)
	WeeklyNotification(weekday time.Weekday, hour, min, sec int, category string)
	MonthlyNotification(day, hour, min, sec int, category string)
}

type AlertDatabase interface {
	CreateAlert(av *entities.AlertVariables) error
	ViewAlert(category string) ([9]interface{}, error)
	RemoveAlert(category string) error
	UpdateAlert(av *entities.AlertVariables) error
}

type Alerts interface {
	AlertSetup(av *entities.AlertVariables) error
	SendAlert(category string) error
	CheckNotification(category string) error
}

type EnvFile interface {
	LoadEnvVariable(key string) string
	WriteEnvFile(av *entities.AuthVariables, dv *entities.DatabaseVariables) error
}

type Email interface {
	SendAlertEmail(category string) error
}

type Conversion interface {
	IntToString(key int) string
	StringToInt(key string) int
}
