package app

import (
	"testing"
)

func TestCreateBudget(t *testing.T) {
	testCases := []struct {
		name        string
		input       *BudgetVariables
		expectedMsg string
	}{
		{
			name: "ValidInput",
			input: &BudgetVariables{
				Category: "TestCategory",
				Amount:   "100",
			},
			expectedMsg: "Created budget for category: Test Category, amount: 100\n",
		},
		{
			name: "EmptyCategory",
			input: &BudgetVariables{
				Category: "",
				Amount:   "100",
			},
			expectedMsg: "Created budget for category: , amount: 100\n",
		},
		{
			name: "EmptyAmount",
			input: &BudgetVariables{
				Category: "CategoryTest",
				Amount:   "",
			},
			expectedMsg: "Created budget for category:Test Category , amount: \n",
		},
	}

	basePath := "../../../db/budget_db/migrations/"
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := CreateBudget(tc.input, basePath)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}
