package app

import (
	"fmt"
	"testing"

	"github.com/ibilalkayy/flow/db/budget_db"
	internal_budget "github.com/ibilalkayy/flow/internal/app/budget"
	"github.com/ibilalkayy/flow/internal/common/functions"
	"github.com/ibilalkayy/flow/internal/entities"
)

type mockDB struct{}

func (m *mockDB) Connection() (*mockRows, error) {
	return &mockRows{}, nil
}

type mockRows struct{}

func (m *mockRows) Close() error {
	return nil
}

func (m *mockRows) Next() bool {
	return true
}

func (m *mockRows) Scan(dest ...interface{}) error {
	// Assuming the budget amount is 100 for testing purposes
	*dest[0].(*string) = "100"
	return nil
}

func (m *mockRows) Err() error {
	return nil
}

func TestCreateBudget(t *testing.T) {
	testCases := []struct {
		name        string
		input       *entities.BudgetVariables
		expectedMsg string
	}{
		{
			name: "ValidInput",
			input: &entities.BudgetVariables{
				Category: "TestCategory",
				Amount:   100,
			},
			expectedMsg: "Created budget for category: TestCategory, amount: 100\n",
		},
		{
			name: "EmptyCategory",
			input: &entities.BudgetVariables{
				Category: "",
				Amount:   100,
			},
			expectedMsg: "Created budget for category: , amount: 100\n",
		},
		{
			name: "EmptyAmount",
			input: &entities.BudgetVariables{
				Category: "CategoryTest",
				Amount:   0,
			},
			expectedMsg: "Created budget for category: TestCategory, amount: \n",
		},
	}

	basePath := "../../../db/budget_db/migrations/"
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := budget_db.CreateBudget(tc.input, basePath)
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
			data, err := budget_db.ViewBudget(tc.category)

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
			err := budget_db.RemoveBudget(tc.category)
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
			amount := functions.StringToInt(tc.amount)
			err := budget_db.UpdateBudget(tc.oldCategory, tc.newCategory, amount)
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
			err := budget_db.GetBudgetData(tc.filepath, tc.filename)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func TestCategoryAmount(t *testing.T) {
	testCases := []struct {
		name        string
		category    string
		expectedMsg string
	}{
		{
			name:        "RightCategory",
			category:    "first",
			expectedMsg: "Found the category amount: 500\n",
		},
		{
			name:        "Wrongcategory",
			category:    "wrong",
			expectedMsg: "Found the category amount: \n",
		},
		{
			name:        "EmptyCategory",
			category:    "",
			expectedMsg: "Found no category amount\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := internal_budget.CategoryAmount(tc.category)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}
