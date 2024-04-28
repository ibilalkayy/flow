package entities

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
