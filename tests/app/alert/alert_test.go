package app

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	internal_alert "github.com/ibilalkayy/flow/internal/app/alert"
	"github.com/ibilalkayy/flow/internal/structs"
)

func TestGenerateUniqueCategory(t *testing.T) {
	// Map to store generated categories to check for uniqueness
	generatedCategories := make(map[string]bool)

	// Regular expression to match the expected format of generated categories
	categoryRegex := regexp.MustCompile(`^total_category_[a-zA-Z0-9]{8}$`)

	// Number of iterations for generating categories
	numIterations := 1000

	for i := 0; i < numIterations; i++ {
		category := "total_cateogory"

		// Check if the generated category matches the expected format
		if !categoryRegex.MatchString(category) {
			t.Errorf("Generated category %q does not match the expected format", category)
		}

		// Check if the generated category is unique
		if _, exists := generatedCategories[category]; exists {
			t.Errorf("Generated category %q is not unique", category)
		}

		// Add the generated category to the map
		generatedCategories[category] = true
	}

	// Print the number of unique categories generated
	fmt.Printf("Generated %d unique categories successfully\n", len(generatedCategories))
}

func TestCreateAlert(t *testing.T) {
	testCases := []struct {
		name        string
		input       *structs.AlertVariables
		expectedMsg string
	}{
		{
			name: "ValidInput",
			input: &structs.AlertVariables{
				Category:  "first",
				Frequency: "hourly",
				Method:    "email",
			},
			expectedMsg: "Created alert for category: first, total: 500, frequency: hourly, method: email\n",
		},
		{
			name: "NoCategory",
			input: &structs.AlertVariables{
				Category:  "",
				Frequency: "hourly",
				Method:    "email",
			},
			expectedMsg: "Created alert for category: , total: 500, frequency: hourly, method: email\n",
		},
		{
			name: "NoTotalAmount",
			input: &structs.AlertVariables{
				Category:  "first",
				Frequency: "hourly",
				Method:    "email",
			},
			expectedMsg: "Created alert for category: first, total: , frequency: hourly, method: email\n",
		},
		{
			name: "NoFrequency",
			input: &structs.AlertVariables{
				Category:  "first",
				Frequency: "",
				Method:    "email",
			},
			expectedMsg: "Created alert for category: first, total: 500, frequency: , method: email\n",
		},
		{
			name: "NoFrequency",
			input: &structs.AlertVariables{
				Category:  "first",
				Frequency: "hourly",
				Method:    "",
			},
			expectedMsg: "Created alert for category: first, total: 500, frequency: hourly, method: \n",
		},
	}

	basePath := "../../../db/budget_db/migrations/"
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := internal_alert.CreateAlert(tc.input, basePath)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func TestAlertSetup(t *testing.T) {
	tests := []struct {
		name        string
		alertVars   *structs.AlertVariables
		expectedErr error
	}{
		// Test cases for valid inputs
		{
			name: "Valid inputs - Total amount alert",
			alertVars: &structs.AlertVariables{
				Frequency: "daily",
				Method:    "email",
			},
			expectedErr: nil,
		},
		{
			name: "Valid inputs - Category amount alert",
			alertVars: &structs.AlertVariables{
				Frequency: "weekly",
				Method:    "cli",
				Category:  "food",
			},
			expectedErr: nil,
		},
		// Test cases for invalid inputs
		{
			name: "Invalid method",
			alertVars: &structs.AlertVariables{
				Frequency: "hourly",
				Method:    "invalid",
			},
			expectedErr: errors.New("invalid alert method"),
		},
		{
			name: "Invalid frequency",
			alertVars: &structs.AlertVariables{
				Frequency: "invalid",
				Method:    "cli",
			},
			expectedErr: errors.New("invalid alert frequency"),
		},
		{
			name: "Missing total or category",
			alertVars: &structs.AlertVariables{
				Frequency: "monthly",
				Method:    "email",
			},
			expectedErr: errors.New("total amount is not given"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := internal_alert.AlertSetup(test.alertVars)
			if err != nil {
				t.Errorf("Expected error: %v, got: %v", test.expectedErr, err)
			}
		})
	}
}

// Mock internal_budget package for testing
type mockBudget struct{}

func (m mockBudget) TotalBudgetAmount() (int, error) {
	return 1000, nil
}

// Mock transaction package for testing
type mockTransaction struct{}

func (m mockTransaction) TransactionAmount() int {
	return 800
}

func TestAlertMessage(t *testing.T) {
	// Mock dependencies
	_ = mockBudget{}
	_ = mockTransaction{}

	// Test case where transaction amount is less than total budget
	err := internal_alert.AlertMessage()
	if err != nil {
		t.Errorf("AlertMessage() returned an error, expected nil: %v", err)
	}

	// Test case where transaction amount is greater than or equal to total budget
	_ = mockTransaction{}
	err = internal_alert.AlertMessage()
	expectedError := "You can't spend more because your budget is set to 1000"
	if err == nil || err.Error() != expectedError {
		t.Errorf("AlertMessage() returned unexpected error: got %v, want %v", err, expectedError)
	}
}
