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
		{1, true, true, "¥1-", nil},
		{10, true, true, "¥10-", nil},
		{100, true, true, "¥100-", nil},
		{1000, true, true, "¥1,000-", nil},
		{10000, true, true, "¥10,000-", nil},
		{100000, true, true, "¥100,000-", nil},
		{1000000, true, true, "¥1,000,000-", nil},
		{10000000, true, true, "¥10,000,000-", nil},
		{100000000, true, true, "¥100,000,000-", nil},
		{1000000000, true, true, "¥1,000,000,000-", nil},
		{10000000000, true, true, "¥10,000,000,000-", nil},
		{100000000000, true, true, "¥100,000,000,000-", nil},
		{1000000000000, true, true, "¥1,000,000,000,000-", nil},
		{10000000000000, true, true, "¥10,000,000,000,000-", nil},
		{100000000000000, true, true, "¥100,000,000,000,000-", nil},
		{1000000000000000, true, true, "¥1,000,000,000,000,000-", nil},
		{9999999999999999, true, true, "¥9,999,999,999,999,999-", nil},
		{10000000000000000, true, true, "", formatter.ErrPriceExceedsLimit},
		{100000, false, false, "100,000", nil},
		{1000000, true, false, "¥1,000,000", nil},
		{10000000000, false, true, "10,000,000,000-", nil},
	}

	for _, test := range tests {
		result, err := formatter.Format(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult || err != test.expectedError {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatInt8(t *testing.T) {
	tests := []struct {
		price          int8
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
		expectedError  error
	}{
		{1, true, true, "¥1-", nil},
		{10, true, true, "¥10-", nil},
		{100, true, true, "¥100-", nil},
		{100, false, false, "100", nil},
		{100, true, false, "¥100", nil},
		{100, false, true, "100-", nil},
	}

	for _, test := range tests {
		result, err := formatter.Format(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult || err != test.expectedError {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}
