package entities

type BudgetVariables struct {
	Category  string
	Amount    int
	Spent     int
	Remaining int
}

type BudgetCalculateVariables struct {
	Category            string
	BudgetAmount        int
	BudgetAmountInDB    int
	SpentAmountInDB     int
	RemainingAmountInDB int
}
