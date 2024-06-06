package interfaces

import (
	"database/sql"
	"time"

	"github.com/ibilalkayy/flow/entities"
)

type Init interface {
	WriteEnvFile(av *entities.AuthVariables, dv *entities.DatabaseVariables) error
}

type Connect interface {
	Connection() (*sql.DB, error)
	Table(filename string, number int) (*sql.DB, error)
	TableExists(tableName string) (bool, error)
}

type TotalAmountDB interface {
	InsertTotalAmount(tv *entities.TotalAmountVariables) error
	ViewTotalAmount() ([5]interface{}, error)
	RemoveTotalAmount(category string) error
	UpdateTotalAmount(tv *entities.TotalAmountVariables) error
	UpdateSpentAndRemaining(spentAmount, remainingAmount int) error
	UpdateStatus(tv *entities.TotalAmountVariables) error
	CalculateRemaining(category string) error
	TotalAmountValues() ([][2]string, [3]interface{}, error)
}

type TotalAmountCategoryDB interface {
	InsertTotalAmountCategory(tv *entities.TotalAmountVariables) error
	ViewTotalAmountCategories() (string, [][2]string, error)
}

type TotalAmount interface {
	SetTotalAmount(totalAmount int, include_category, label string) error
	HandleExistingTables(totalAmount int, tav, tacv entities.TotalAmountVariables) error
	HandleMissingTables(tav, tacv entities.TotalAmountVariables) error
}

type SpendAmount interface {
	UpdateBudgetAndTotalAmount(sv *entities.SpendingVariables) error
	HandleExceededBudget(sv *entities.SpendingVariables) error
	ValidBudgetValues(sv *entities.SpendingVariables) error
	SpendMoney(category string, spending_amount int) error
	StoreHistory(category string, spending_amount int) error

	HourlyNotification(category string)
	DailyNotification(hour, min int, category string)
	WeeklyNotification(weekday time.Weekday, hour, min int, category string)
	MonthlyNotification(day, hour, min int, category string)
}

type ManageBudget interface {
	CreateBudget(bv *entities.BudgetVariables) error
	TakeBudgetAmount() ([]string, []int, error)
	ListOfExpection(bv *entities.BudgetVariables) ([]int, []int, error)
	ViewBudget(category string) ([5]interface{}, error)
	UpdateBudget(bv *entities.BudgetVariables, new_category string) error
	UpdateBudgetCategory(new, old string) error
	RemoveBudget(category string) error
	AddBudgetExpenditure(spent int, category string) error
	GetBudgetData(filepath, filename string) error
	CalculateRemaining(cr *entities.BudgetCalculateVariables) ([2]int, error)

	InsertHistory(hv *entities.HistoryVariables) error
	ViewHistory(category string) ([2]interface{}, error)
	RemoveHistory(category string) error
}

type Budget interface {
	CategoryAmount(category string) (string, int, error)
}

type HandleEmail interface {
	SendAlertEmail(category string) error
}

type ManageAlerts interface {
	AlertSetup(av *entities.AlertVariables) error
	SendAlert(category string) error
	SendNotification(category string) error
	WriteNotificationValues(av *entities.AlertVariables) error
}

type AlertDB interface {
	CreateAlert(av *entities.AlertVariables) error
	ViewAlert(category string) ([8]interface{}, error)
	RemoveAlert(category string) error
	UpdateAlert(av *entities.AlertVariables) error
}

type Env interface {
	LoadEnvVariable(key string) string
}

type Common interface {
	IntToString(key int) string
	StringToInt(key string) int
}

type Dependencies struct {
	Init                Init
	Connect             Connect
	TotalAmount         TotalAmountDB
	TotalAmountCategory TotalAmountCategoryDB
	Total               TotalAmount
	SpendAmount         SpendAmount
	ManageBudget        ManageBudget
	HandleEmail         HandleEmail
	ManageAlerts        ManageAlerts
	AlertDB             AlertDB
	Budget              Budget
	Env                 Env
	Common              Common
}
