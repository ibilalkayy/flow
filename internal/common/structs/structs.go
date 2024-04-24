package structs

type AuthVariables struct {
	Username    string
	Gmail       string
	AppPassword string
}

type DatabaseVariables struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type BudgetVariables struct {
	Category  string
	Amount    int
	Spent     int
	Remaining int
}

type HistoryVariables struct {
	Date          string
	Time          string
	Category      string
	Amount        int
	TransactionID string
	Blockchain    string
	Address       string
}

type SpendingVariables struct {
	Category                      string
	CategoryName                  string
	TotalAmountStatus             string
	IncludedCatogeries            [][2]string
	TotalAmount                   int
	SpendingAmount                int
	TotalAmountSpent              int
	BudgetCategoryAmount          int
	BudgetCategorySpentAmount     int
	BudgetCategoryRemainingAmount int
}

type AlertVariables struct {
	Category  string
	Method    string
	Frequency string
	Days      int
	Weekdays  string
	Hours     int
	Minutes   int
	Seconds   int
}

type EmailVariables struct {
	Username       string
	Category       string
	CategoryAmount int
}

type TotalAmountVariables struct {
	TotalAmount     int
	SpentAmount     int
	RemainingAmount int
	Included        string
	NewCategory     string
	Label           string
	Status          string
}
