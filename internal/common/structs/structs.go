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

type AlertVariables struct {
	Category  string
	Frequency string
	Method    string
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
	Amount   int
	Included string
	Excluded string
	Label    string
}
