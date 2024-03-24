package structs

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
