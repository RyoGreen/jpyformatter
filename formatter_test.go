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
		{-100, true, true, "¥-100-", nil},
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
		{-100, true, true, "¥-100-", nil},
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

func TestFormatInt16(t *testing.T) {
	tests := []struct {
		price          int16
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
		{-100, true, true, "¥-100-", nil},
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

func TestFormatInt32(t *testing.T) {
	tests := []struct {
		price          int32
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
		{100000, false, false, "100,000", nil},
		{1000000, true, false, "¥1,000,000", nil},
		{1000000, false, true, "1,000,000-", nil},
		{-100, true, true, "¥-100-", nil},
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

func TestFormatInt64(t *testing.T) {
	tests := []struct {
		price          int64
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
		{-100, true, true, "¥-100-", nil},
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

func TestFormatUInt(t *testing.T) {
	tests := []struct {
		price          uint
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

func TestFormatUInt8(t *testing.T) {
	tests := []struct {
		price          uint8
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

func TestFormatUInt16(t *testing.T) {
	tests := []struct {
		price          uint16
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

func TestFormatUInt32(t *testing.T) {
	tests := []struct {
		price          uint32
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
		{100000, false, false, "100,000", nil},
		{1000000, true, false, "¥1,000,000", nil},
		{1000000, false, true, "1,000,000-", nil},
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

func TestFormatUInt64(t *testing.T) {
	tests := []struct {
		price          uint64
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

func TestFormatUIntptr(t *testing.T) {
	tests := []struct {
		price          uintptr
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

func TestFormatArbitraryType(t *testing.T) {
	type MyInt int
	tests := []struct {
		price          MyInt
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
