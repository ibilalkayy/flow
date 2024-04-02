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
	Spent    string
}

type AlertVariables struct {
	Category       string
	CategoryAmount string
	Frequency      string
	Method         string
	Days           string
	Weekdays       string
	Hours          string
	Minutes        string
	Seconds        string
}

type EmailVariables struct {
	Username       string
	Category       string
	CategoryAmount string
}
