package entities

type BudgetVariables struct {
	Category  string
	Amount    int
	Spent     int
	Remaining int
}

type BudgetCalculateVariables struct {
	BudgetAmount        int
	BudgetAmountInDB    int
	SpentAmountInDB     int
	RemainingAmountInDB int
}
