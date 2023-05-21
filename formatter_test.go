package formatter_test

import (
	"fmt"
	"testing"

	formatter "github.com/RyoGreen/jpyformatter"
)

func TestFormatInt(t *testing.T) {
	tests := []struct {
		price          int
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
		expectedError  error
	}{
		{100000, false, false, "100,000", nil},
		{1000000, true, true, "¥1,000,000-", nil},
		{10000000000, true, false, "¥10,000,000,000", nil},
		{10000000000000000, true, true, "", formatter.ErrPriceExceedsLimit},
	}

	for _, test := range tests {
		result, err := formatter.Format(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult || err != test.expectedError {
			fmt.Printf("Test failed. Price: %d, Prefix: %v, Suffix: %v\n", test.price, test.prefixEnabled, test.suffixEnabled)
			fmt.Printf("Expected result: %s, Expected error: %v\n", test.expectedResult, test.expectedError)
			fmt.Printf("Actual result: %s, Actual error: %v\n", result, err)
		} else {
			fmt.Println("Test passed.")
		}
	}
}
