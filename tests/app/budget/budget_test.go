package app

import (
	"fmt"
	"testing"

	internal_budget "github.com/ibilalkayy/flow/internal/app/budget"
	"github.com/ibilalkayy/flow/internal/structs"
)

func TestCreateBudget(t *testing.T) {
	testCases := []struct {
		name        string
		input       *structs.BudgetVariables
		expectedMsg string
	}{
		{
			name: "ValidInput",
			input: &structs.BudgetVariables{
				Category: "TestCategory",
				Amount:   "100",
			},
			expectedMsg: "Created budget for category: TestCategory, amount: 100\n",
		},
		{
			name: "EmptyCategory",
			input: &structs.BudgetVariables{
				Category: "",
				Amount:   "100",
			},
			expectedMsg: "Created budget for category: , amount: 100\n",
		},
		{
			name: "EmptyAmount",
			input: &structs.BudgetVariables{
				Category: "CategoryTest",
				Amount:   "",
			},
			expectedMsg: "Created budget for category: TestCategory, amount: \n",
		},
	}

	basePath := "../../../db/budget_db/migrations/"
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := internal_budget.CreateBudget(tc.input, basePath)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func TestViewBudget(t *testing.T) {
	testCases := []struct {
		name        string
		category    string
		expectedMsg string
	}{
		{
			name:        "FilledCategory",
			category:    "TestCategory",
			expectedMsg: "Viewed the Budget info of category: TestCategory",
		},
		{
			name:        "EmptyCategory",
			category:    "",
			expectedMsg: "Viewed the Budget info of category: ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := internal_budget.ViewBudget(tc.category)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			fmt.Println(data)
		})
	}
}

func TestRemoveBudget(t *testing.T) {
	testCases := []struct {
		name        string
		category    string
		expectedMsg string
	}{
		{
			name:        "FilledCategory",
			category:    "TestCategory",
			expectedMsg: "Removed the budget info of category: TestCategory\n",
		},
		{
			name:        "EmptyCategory",
			category:    "",
			expectedMsg: "Removed the budget info of category: TestCategory\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := internal_budget.RemoveBudget(tc.category)
			if err != nil {
				t.Errorf("Expected no error: got %v", err)
			}
		})
	}
}

func TestUpdateBudget(t *testing.T) {
	testCases := []struct {
		name        string
		oldCategory string
		newCategory string
		amount      string
		expectedMsg string
	}{
		{
			name:        "CategoryAndAmount",
			oldCategory: "TestCategory",
			newCategory: "NewCategory",
			amount:      "200",
			expectedMsg: "Updated the category: NewCategory, amount: 200\n",
		},
		{
			name:        "Category",
			oldCategory: "NewCategory",
			newCategory: "SecondNewCategory",
			amount:      "",
			expectedMsg: "Updated the category: NewCategory, amount: \n",
		},
		{
			name:        "Amount",
			oldCategory: "SecondNewCategory",
			newCategory: "",
			amount:      "150",
			expectedMsg: "Updated the category: , amount: 100\n",
		},
		{
			name:        "EmptyBudget",
			oldCategory: "SecondNewCategory",
			newCategory: "",
			amount:      "",
			expectedMsg: "Updated the category: , amount: \n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := internal_budget.UpdateBudget(tc.oldCategory, tc.newCategory, tc.amount)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func TestGetBudgetData(t *testing.T) {
	testCases := []struct {
		name        string
		filepath    string
		filename    string
		expectedMsg string
	}{
		{
			name:        "FilenameAndPath",
			filepath:    "/mnt/d/go/src/github.com/ibilalkayy/flow/tests/app/budget",
			filename:    "hello.csv",
			expectedMsg: "Got the file by giving filename: hello.csv, filepath: /mnt/d/go/src/github.com/ibilalkayy/flow/tests/app/budge\n",
		},
		{
			name:        "Filename",
			filepath:    "",
			filename:    "hello.csv",
			expectedMsg: "Got no file by giving filename: hello.csv, filepath: \n",
		},
		{
			name:        "Filepath",
			filepath:    "/mnt/d/go/src/github.com/ibilalkayy/flow/tests/app/budget",
			filename:    "",
			expectedMsg: "Got no file by giving filename: , filepath: /mnt/d/go/src/github.com/ibilalkayy/flow/tests/app/budget\n",
		},
		{
			name:        "EmptyPaths",
			filepath:    "",
			filename:    "",
			expectedMsg: "Got no file by giving filename: , filepath: \n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := internal_budget.GetBudgetData(tc.filepath, tc.filename)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}
