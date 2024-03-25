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
	Category string
	Amount   string
}

type AlertVariables struct {
	Category       string
	CategoryAmount string
	Total          string
	Frequency      string
	Method         string
}

type EmailVariables struct {
	Username       string
	Category       string
	CategoryAmount string
	TotalAmount    string
}
