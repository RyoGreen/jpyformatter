package formatter_test

import (
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
		{1000, true, false, "¥1,000", nil},
		{1000000, false, true, "1,000,000-", nil},
		{1000000, false, false, "1,000,000", nil},
		{999999999999, true, true, "¥999,999,999,999-", nil},
		{999999999999, true, false, "¥999,999,999,999", nil},
		{999999999999, false, true, "999,999,999,999-", nil},
		{999999999999, false, false, "999,999,999,999", nil},
		{1000000000000, true, true, "", formatter.ErrPriceExceedsLimit},
		{1000000000000, true, false, "", formatter.ErrPriceExceedsLimit},
		{1000000000000, false, true, "", formatter.ErrPriceExceedsLimit},
		{1000000000000, false, false, "", formatter.ErrPriceExceedsLimit},
	}

	for _, test := range tests {
		result, err := formatter.Format(test.price, test.prefixEnabled, test.suffixEnabled)

		if result != test.expectedResult {
			t.Errorf("Price: %v, PrefixEnabled: %v, SuffixEnabled: %v - Expected result: %s, Got: %s", test.price, test.prefixEnabled, test.suffixEnabled, test.expectedResult, result)
		}

		if (err != nil && test.expectedError == nil) || (err == nil && test.expectedError != nil) || (err != nil && test.expectedError != nil && err.Error() != test.expectedError.Error()) {
			t.Errorf("Price: %v, PrefixEnabled: %v, SuffixEnabled: %v - Expected error: %v, Got: %v", test.price, test.prefixEnabled, test.suffixEnabled, test.expectedError, err)
		}
	}
}
