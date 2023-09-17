package formatter_test

import (
	"fmt"
	"testing"

	formatter "github.com/RyoGreen/jpyformatter"
)

func TestFormatXInt(t *testing.T) {
	tests := []struct {
		price          int
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1000, true, true, "¥1,000-"},
		{10000, true, true, "¥10,000-"},
		{100000, true, true, "¥100,000-"},
		{1000000, true, true, "¥1,000,000-"},
		{10000000, true, true, "¥10,000,000-"},
		{100000000, true, true, "¥100,000,000-"},
		{1000000000, true, true, "¥1,000,000,000-"},
		{10000000000, true, true, "¥10,000,000,000-"},
		{100000000000, true, true, "¥100,000,000,000-"},
		{1000000000000, true, true, "¥1,000,000,000,000-"},
		{10000000000000, true, true, "¥10,000,000,000,000-"},
		{99999999999999, true, true, "¥99,999,999,999,999-"},
		{100000000000000, true, true, "¥0-"},
		{-1, true, true, "¥-1-"},
		{-10, true, true, "¥-10-"},
		{-100, true, true, "¥-100-"},
		{-1000, true, true, "¥-1,000-"},
		{-10000, true, true, "¥-10,000-"},
		{-100000, true, true, "¥-100,000-"},
		{-1000000, true, true, "¥-1,000,000-"},
		{-10000000, true, true, "¥-10,000,000-"},
		{-100000000, true, true, "¥-100,000,000-"},
		{-1000000000, true, true, "¥-1,000,000,000-"},
		{-10000000000, true, true, "¥-10,000,000,000-"},
		{-100000000000, true, true, "¥-100,000,000,000-"},
		{-1000000000000, true, true, "¥-1,000,000,000,000-"},
		{-10000000000000, true, true, "¥-10,000,000,000,000-"},
		{-99999999999999, true, true, "¥-99,999,999,999,999-"},
		{-100000000000000, true, true, "¥0-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1000, false, true, "1,000-"},
		{10000, false, true, "10,000-"},
		{100000, false, true, "100,000-"},
		{1000000, false, true, "1,000,000-"},
		{10000000, false, true, "10,000,000-"},
		{100000000, false, true, "100,000,000-"},
		{1000000000, false, true, "1,000,000,000-"},
		{10000000000, false, true, "10,000,000,000-"},
		{100000000000, false, true, "100,000,000,000-"},
		{1000000000000, false, true, "1,000,000,000,000-"},
		{10000000000000, false, true, "10,000,000,000,000-"},
		{99999999999999, false, true, "99,999,999,999,999-"},
		{100000000000000, false, true, "0-"},
		{-1, false, true, "-1-"},
		{-10, false, true, "-10-"},
		{-100, false, true, "-100-"},
		{-1000, false, true, "-1,000-"},
		{-10000, false, true, "-10,000-"},
		{-100000, false, true, "-100,000-"},
		{-1000000, false, true, "-1,000,000-"},
		{-10000000, false, true, "-10,000,000-"},
		{-100000000, false, true, "-100,000,000-"},
		{-1000000000, false, true, "-1,000,000,000-"},
		{-10000000000, false, true, "-10,000,000,000-"},
		{-100000000000, false, true, "-100,000,000,000-"},
		{-1000000000000, false, true, "-1,000,000,000,000-"},
		{-10000000000000, false, true, "-10,000,000,000,000-"},
		{-99999999999999, false, true, "-99,999,999,999,999-"},
		{-100000000000000, false, true, "0-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1000, true, false, "¥1,000"},
		{10000, true, false, "¥10,000"},
		{100000, true, false, "¥100,000"},
		{1000000, true, false, "¥1,000,000"},
		{10000000, true, false, "¥10,000,000"},
		{100000000, true, false, "¥100,000,000"},
		{1000000000, true, false, "¥1,000,000,000"},
		{10000000000, true, false, "¥10,000,000,000"},
		{100000000000, true, false, "¥100,000,000,000"},
		{1000000000000, true, false, "¥1,000,000,000,000"},
		{10000000000000, true, false, "¥10,000,000,000,000"},
		{99999999999999, true, false, "¥99,999,999,999,999"},
		{100000000000000, true, false, "¥0"},
		{-1, true, false, "¥-1"},
		{-10, true, false, "¥-10"},
		{-100, true, false, "¥-100"},
		{-1000, true, false, "¥-1,000"},
		{-10000, true, false, "¥-10,000"},
		{-100000, true, false, "¥-100,000"},
		{-1000000, true, false, "¥-1,000,000"},
		{-10000000, true, false, "¥-10,000,000"},
		{-100000000, true, false, "¥-100,000,000"},
		{-1000000000, true, false, "¥-1,000,000,000"},
		{-10000000000, true, false, "¥-10,000,000,000"},
		{-100000000000, true, false, "¥-100,000,000,000"},
		{-1000000000000, true, false, "¥-1,000,000,000,000"},
		{-10000000000000, true, false, "¥-10,000,000,000,000"},
		{-99999999999999, true, false, "¥-99,999,999,999,999"},
		{-100000000000000, true, false, "¥0"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{1000, false, false, "1,000"},
		{10000, false, false, "10,000"},
		{100000, false, false, "100,000"},
		{1000000, false, false, "1,000,000"},
		{10000000, false, false, "10,000,000"},
		{100000000, false, false, "100,000,000"},
		{1000000000, false, false, "1,000,000,000"},
		{10000000000, false, false, "10,000,000,000"},
		{100000000000, false, false, "100,000,000,000"},
		{1000000000000, false, false, "1,000,000,000,000"},
		{10000000000000, false, false, "10,000,000,000,000"},
		{99999999999999, false, false, "99,999,999,999,999"},
		{100000000000000, false, false, "0"},
		{-1, false, false, "-1"},
		{-10, false, false, "-10"},
		{-100, false, false, "-100"},
		{-1000, false, false, "-1,000"},
		{-10000, false, false, "-10,000"},
		{-100000, false, false, "-100,000"},
		{-1000000, false, false, "-1,000,000"},
		{-10000000, false, false, "-10,000,000"},
		{-100000000, false, false, "-100,000,000"},
		{-1000000000, false, false, "-1,000,000,000"},
		{-10000000000, false, false, "-10,000,000,000"},
		{-100000000000, false, false, "-100,000,000,000"},
		{-1000000000000, false, false, "-1,000,000,000,000"},
		{-10000000000000, false, false, "-10,000,000,000,000"},
		{-99999999999999, false, false, "-99,999,999,999,999"},
		{-100000000000000, false, false, "0"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXInt8(t *testing.T) {
	tests := []struct {
		price          int8
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{-1, true, true, "¥-1-"},
		{-10, true, true, "¥-10-"},
		{-100, true, true, "¥-100-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{-1, false, true, "-1-"},
		{-10, false, true, "-10-"},
		{-100, false, true, "-100-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{-1, true, false, "¥-1"},
		{-10, true, false, "¥-10"},
		{-100, true, false, "¥-100"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{-1, false, false, "-1"},
		{-10, false, false, "-10"},
		{-100, false, false, "-100"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXInt16(t *testing.T) {
	tests := []struct {
		price          int16
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1000, true, true, "¥1,000-"},
		{10000, true, true, "¥10,000-"},
		{-1, true, true, "¥-1-"},
		{-10, true, true, "¥-10-"},
		{-100, true, true, "¥-100-"},
		{-1000, true, true, "¥-1,000-"},
		{-10000, true, true, "¥-10,000-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1000, false, true, "1,000-"},
		{10000, false, true, "10,000-"},
		{-1, false, true, "-1-"},
		{-10, false, true, "-10-"},
		{-100, false, true, "-100-"},
		{-1000, false, true, "-1,000-"},
		{-10000, false, true, "-10,000-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1000, true, false, "¥1,000"},
		{10000, true, false, "¥10,000"},
		{-1, true, false, "¥-1"},
		{-10, true, false, "¥-10"},
		{-100, true, false, "¥-100"},
		{-1000, true, false, "¥-1,000"},
		{-10000, true, false, "¥-10,000"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{1000, false, false, "1,000"},
		{10000, false, false, "10,000"},
		{-1, false, false, "-1"},
		{-10, false, false, "-10"},
		{-100, false, false, "-100"},
		{-1000, false, false, "-1,000"},
		{-10000, false, false, "-10,000"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXInt32(t *testing.T) {
	tests := []struct {
		price          int32
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1000, true, true, "¥1,000-"},
		{10000, true, true, "¥10,000-"},
		{100000, true, true, "¥100,000-"},
		{1000000, true, true, "¥1,000,000-"},
		{10000000, true, true, "¥10,000,000-"},
		{100000000, true, true, "¥100,000,000-"},
		{1000000000, true, true, "¥1,000,000,000-"},
		{-1, true, true, "¥-1-"},
		{-10, true, true, "¥-10-"},
		{-100, true, true, "¥-100-"},
		{-1000, true, true, "¥-1,000-"},
		{-10000, true, true, "¥-10,000-"},
		{-100000, true, true, "¥-100,000-"},
		{-1000000, true, true, "¥-1,000,000-"},
		{-10000000, true, true, "¥-10,000,000-"},
		{-100000000, true, true, "¥-100,000,000-"},
		{-1000000000, true, true, "¥-1,000,000,000-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1000, false, true, "1,000-"},
		{10000, false, true, "10,000-"},
		{100000, false, true, "100,000-"},
		{1000000, false, true, "1,000,000-"},
		{10000000, false, true, "10,000,000-"},
		{100000000, false, true, "100,000,000-"},
		{1000000000, false, true, "1,000,000,000-"},
		{-1, false, true, "-1-"},
		{-10, false, true, "-10-"},
		{-100, false, true, "-100-"},
		{-1000, false, true, "-1,000-"},
		{-10000, false, true, "-10,000-"},
		{-100000, false, true, "-100,000-"},
		{-1000000, false, true, "-1,000,000-"},
		{-10000000, false, true, "-10,000,000-"},
		{-100000000, false, true, "-100,000,000-"},
		{-1000000000, false, true, "-1,000,000,000-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1000, true, false, "¥1,000"},
		{10000, true, false, "¥10,000"},
		{100000, true, false, "¥100,000"},
		{1000000, true, false, "¥1,000,000"},
		{10000000, true, false, "¥10,000,000"},
		{100000000, true, false, "¥100,000,000"},
		{1000000000, true, false, "¥1,000,000,000"},
		{-1, true, false, "¥-1"},
		{-10, true, false, "¥-10"},
		{-100, true, false, "¥-100"},
		{-1000, true, false, "¥-1,000"},
		{-10000, true, false, "¥-10,000"},
		{-100000, true, false, "¥-100,000"},
		{-1000000, true, false, "¥-1,000,000"},
		{-10000000, true, false, "¥-10,000,000"},
		{-100000000, true, false, "¥-100,000,000"},
		{-1000000000, true, false, "¥-1,000,000,000"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{1000, false, false, "1,000"},
		{10000, false, false, "10,000"},
		{100000, false, false, "100,000"},
		{1000000, false, false, "1,000,000"},
		{10000000, false, false, "10,000,000"},
		{100000000, false, false, "100,000,000"},
		{1000000000, false, false, "1,000,000,000"},
		{-1, false, false, "-1"},
		{-10, false, false, "-10"},
		{-100, false, false, "-100"},
		{-1000, false, false, "-1,000"},
		{-10000, false, false, "-10,000"},
		{-100000, false, false, "-100,000"},
		{-1000000, false, false, "-1,000,000"},
		{-10000000, false, false, "-10,000,000"},
		{-100000000, false, false, "-100,000,000"},
		{-1000000000, false, false, "-1,000,000,000"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXInt64(t *testing.T) {
	tests := []struct {
		price          int64
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1000, true, true, "¥1,000-"},
		{10000, true, true, "¥10,000-"},
		{100000, true, true, "¥100,000-"},
		{1000000, true, true, "¥1,000,000-"},
		{10000000, true, true, "¥10,000,000-"},
		{100000000, true, true, "¥100,000,000-"},
		{1000000000, true, true, "¥1,000,000,000-"},
		{10000000000, true, true, "¥10,000,000,000-"},
		{100000000000, true, true, "¥100,000,000,000-"},
		{1000000000000, true, true, "¥1,000,000,000,000-"},
		{10000000000000, true, true, "¥10,000,000,000,000-"},
		{99999999999999, true, true, "¥99,999,999,999,999-"},
		{100000000000000, true, true, "¥0-"},
		{-1, true, true, "¥-1-"},
		{-10, true, true, "¥-10-"},
		{-100, true, true, "¥-100-"},
		{-1000, true, true, "¥-1,000-"},
		{-10000, true, true, "¥-10,000-"},
		{-100000, true, true, "¥-100,000-"},
		{-1000000, true, true, "¥-1,000,000-"},
		{-10000000, true, true, "¥-10,000,000-"},
		{-100000000, true, true, "¥-100,000,000-"},
		{-1000000000, true, true, "¥-1,000,000,000-"},
		{-10000000000, true, true, "¥-10,000,000,000-"},
		{-100000000000, true, true, "¥-100,000,000,000-"},
		{-1000000000000, true, true, "¥-1,000,000,000,000-"},
		{-10000000000000, true, true, "¥-10,000,000,000,000-"},
		{-99999999999999, true, true, "¥-99,999,999,999,999-"},
		{-100000000000000, true, true, "¥0-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1000, false, true, "1,000-"},
		{10000, false, true, "10,000-"},
		{100000, false, true, "100,000-"},
		{1000000, false, true, "1,000,000-"},
		{10000000, false, true, "10,000,000-"},
		{100000000, false, true, "100,000,000-"},
		{1000000000, false, true, "1,000,000,000-"},
		{10000000000, false, true, "10,000,000,000-"},
		{100000000000, false, true, "100,000,000,000-"},
		{1000000000000, false, true, "1,000,000,000,000-"},
		{10000000000000, false, true, "10,000,000,000,000-"},
		{99999999999999, false, true, "99,999,999,999,999-"},
		{100000000000000, false, true, "0-"},
		{-1, false, true, "-1-"},
		{-10, false, true, "-10-"},
		{-100, false, true, "-100-"},
		{-1000, false, true, "-1,000-"},
		{-10000, false, true, "-10,000-"},
		{-100000, false, true, "-100,000-"},
		{-1000000, false, true, "-1,000,000-"},
		{-10000000, false, true, "-10,000,000-"},
		{-100000000, false, true, "-100,000,000-"},
		{-1000000000, false, true, "-1,000,000,000-"},
		{-10000000000, false, true, "-10,000,000,000-"},
		{-100000000000, false, true, "-100,000,000,000-"},
		{-1000000000000, false, true, "-1,000,000,000,000-"},
		{-10000000000000, false, true, "-10,000,000,000,000-"},
		{-99999999999999, false, true, "-99,999,999,999,999-"},
		{-100000000000000, false, true, "0-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1000, true, false, "¥1,000"},
		{10000, true, false, "¥10,000"},
		{100000, true, false, "¥100,000"},
		{1000000, true, false, "¥1,000,000"},
		{10000000, true, false, "¥10,000,000"},
		{100000000, true, false, "¥100,000,000"},
		{1000000000, true, false, "¥1,000,000,000"},
		{10000000000, true, false, "¥10,000,000,000"},
		{100000000000, true, false, "¥100,000,000,000"},
		{1000000000000, true, false, "¥1,000,000,000,000"},
		{10000000000000, true, false, "¥10,000,000,000,000"},
		{99999999999999, true, false, "¥99,999,999,999,999"},
		{100000000000000, true, false, "¥0"},
		{-1, true, false, "¥-1"},
		{-10, true, false, "¥-10"},
		{-100, true, false, "¥-100"},
		{-1000, true, false, "¥-1,000"},
		{-10000, true, false, "¥-10,000"},
		{-100000, true, false, "¥-100,000"},
		{-1000000, true, false, "¥-1,000,000"},
		{-10000000, true, false, "¥-10,000,000"},
		{-100000000, true, false, "¥-100,000,000"},
		{-1000000000, true, false, "¥-1,000,000,000"},
		{-10000000000, true, false, "¥-10,000,000,000"},
		{-100000000000, true, false, "¥-100,000,000,000"},
		{-1000000000000, true, false, "¥-1,000,000,000,000"},
		{-10000000000000, true, false, "¥-10,000,000,000,000"},
		{-99999999999999, true, false, "¥-99,999,999,999,999"},
		{-100000000000000, true, false, "¥0"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{1000, false, false, "1,000"},
		{10000, false, false, "10,000"},
		{100000, false, false, "100,000"},
		{1000000, false, false, "1,000,000"},
		{10000000, false, false, "10,000,000"},
		{100000000, false, false, "100,000,000"},
		{1000000000, false, false, "1,000,000,000"},
		{10000000000, false, false, "10,000,000,000"},
		{100000000000, false, false, "100,000,000,000"},
		{1000000000000, false, false, "1,000,000,000,000"},
		{10000000000000, false, false, "10,000,000,000,000"},
		{99999999999999, false, false, "99,999,999,999,999"},
		{100000000000000, false, false, "0"},
		{-1, false, false, "-1"},
		{-10, false, false, "-10"},
		{-100, false, false, "-100"},
		{-1000, false, false, "-1,000"},
		{-10000, false, false, "-10,000"},
		{-100000, false, false, "-100,000"},
		{-1000000, false, false, "-1,000,000"},
		{-10000000, false, false, "-10,000,000"},
		{-100000000, false, false, "-100,000,000"},
		{-1000000000, false, false, "-1,000,000,000"},
		{-10000000000, false, false, "-10,000,000,000"},
		{-100000000000, false, false, "-100,000,000,000"},
		{-1000000000000, false, false, "-1,000,000,000,000"},
		{-10000000000000, false, false, "-10,000,000,000,000"},
		{-99999999999999, false, false, "-99,999,999,999,999"},
		{-100000000000000, false, false, "0"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXUInt(t *testing.T) {
	tests := []struct {
		price          uint
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1000, true, true, "¥1,000-"},
		{10000, true, true, "¥10,000-"},
		{100000, true, true, "¥100,000-"},
		{1000000, true, true, "¥1,000,000-"},
		{10000000, true, true, "¥10,000,000-"},
		{100000000, true, true, "¥100,000,000-"},
		{1000000000, true, true, "¥1,000,000,000-"},
		{10000000000, true, true, "¥10,000,000,000-"},
		{100000000000, true, true, "¥100,000,000,000-"},
		{1000000000000, true, true, "¥1,000,000,000,000-"},
		{10000000000000, true, true, "¥10,000,000,000,000-"},
		{99999999999999, true, true, "¥99,999,999,999,999-"},
		{100000000000000, true, true, "¥0-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1000, false, true, "1,000-"},
		{10000, false, true, "10,000-"},
		{100000, false, true, "100,000-"},
		{1000000, false, true, "1,000,000-"},
		{10000000, false, true, "10,000,000-"},
		{100000000, false, true, "100,000,000-"},
		{1000000000, false, true, "1,000,000,000-"},
		{10000000000, false, true, "10,000,000,000-"},
		{100000000000, false, true, "100,000,000,000-"},
		{1000000000000, false, true, "1,000,000,000,000-"},
		{10000000000000, false, true, "10,000,000,000,000-"},
		{99999999999999, false, true, "99,999,999,999,999-"},
		{100000000000000, false, true, "0-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1000, true, false, "¥1,000"},
		{10000, true, false, "¥10,000"},
		{100000, true, false, "¥100,000"},
		{1000000, true, false, "¥1,000,000"},
		{10000000, true, false, "¥10,000,000"},
		{100000000, true, false, "¥100,000,000"},
		{1000000000, true, false, "¥1,000,000,000"},
		{10000000000, true, false, "¥10,000,000,000"},
		{100000000000, true, false, "¥100,000,000,000"},
		{1000000000000, true, false, "¥1,000,000,000,000"},
		{10000000000000, true, false, "¥10,000,000,000,000"},
		{99999999999999, true, false, "¥99,999,999,999,999"},
		{100000000000000, true, false, "¥0"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{1000, false, false, "1,000"},
		{10000, false, false, "10,000"},
		{100000, false, false, "100,000"},
		{1000000, false, false, "1,000,000"},
		{10000000, false, false, "10,000,000"},
		{100000000, false, false, "100,000,000"},
		{1000000000, false, false, "1,000,000,000"},
		{10000000000, false, false, "10,000,000,000"},
		{100000000000, false, false, "100,000,000,000"},
		{1000000000000, false, false, "1,000,000,000,000"},
		{10000000000000, false, false, "10,000,000,000,000"},
		{99999999999999, false, false, "99,999,999,999,999"},
		{100000000000000, false, false, "0"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXUInt8(t *testing.T) {
	tests := []struct {
		price          uint8
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXUInt16(t *testing.T) {
	tests := []struct {
		price          uint16
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1000, true, true, "¥1,000-"},
		{10000, true, true, "¥10,000-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1000, false, true, "1,000-"},
		{10000, false, true, "10,000-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1000, true, false, "¥1,000"},
		{10000, true, false, "¥10,000"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{1000, false, false, "1,000"},
		{10000, false, false, "10,000"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXUInt32(t *testing.T) {
	tests := []struct {
		price          uint32
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1000, true, true, "¥1,000-"},
		{10000, true, true, "¥10,000-"},
		{100000, true, true, "¥100,000-"},
		{1000000, true, true, "¥1,000,000-"},
		{10000000, true, true, "¥10,000,000-"},
		{100000000, true, true, "¥100,000,000-"},
		{1000000000, true, true, "¥1,000,000,000-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1000, false, true, "1,000-"},
		{10000, false, true, "10,000-"},
		{100000, false, true, "100,000-"},
		{1000000, false, true, "1,000,000-"},
		{10000000, false, true, "10,000,000-"},
		{100000000, false, true, "100,000,000-"},
		{1000000000, false, true, "1,000,000,000-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1000, true, false, "¥1,000"},
		{10000, true, false, "¥10,000"},
		{100000, true, false, "¥100,000"},
		{1000000, true, false, "¥1,000,000"},
		{10000000, true, false, "¥10,000,000"},
		{100000000, true, false, "¥100,000,000"},
		{1000000000, true, false, "¥1,000,000,000"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{1000, false, false, "1,000"},
		{10000, false, false, "10,000"},
		{100000, false, false, "100,000"},
		{1000000, false, false, "1,000,000"},
		{10000000, false, false, "10,000,000"},
		{100000000, false, false, "100,000,000"},
		{1000000000, false, false, "1,000,000,000"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXUInt64(t *testing.T) {
	tests := []struct {
		price          uint64
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1000, true, true, "¥1,000-"},
		{10000, true, true, "¥10,000-"},
		{100000, true, true, "¥100,000-"},
		{1000000, true, true, "¥1,000,000-"},
		{10000000, true, true, "¥10,000,000-"},
		{100000000, true, true, "¥100,000,000-"},
		{1000000000, true, true, "¥1,000,000,000-"},
		{10000000000, true, true, "¥10,000,000,000-"},
		{100000000000, true, true, "¥100,000,000,000-"},
		{1000000000000, true, true, "¥1,000,000,000,000-"},
		{10000000000000, true, true, "¥10,000,000,000,000-"},
		{99999999999999, true, true, "¥99,999,999,999,999-"},
		{100000000000000, true, true, "¥0-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1000, false, true, "1,000-"},
		{10000, false, true, "10,000-"},
		{100000, false, true, "100,000-"},
		{1000000, false, true, "1,000,000-"},
		{10000000, false, true, "10,000,000-"},
		{100000000, false, true, "100,000,000-"},
		{1000000000, false, true, "1,000,000,000-"},
		{10000000000, false, true, "10,000,000,000-"},
		{100000000000, false, true, "100,000,000,000-"},
		{1000000000000, false, true, "1,000,000,000,000-"},
		{10000000000000, false, true, "10,000,000,000,000-"},
		{99999999999999, false, true, "99,999,999,999,999-"},
		{100000000000000, false, true, "0-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1000, true, false, "¥1,000"},
		{10000, true, false, "¥10,000"},
		{100000, true, false, "¥100,000"},
		{1000000, true, false, "¥1,000,000"},
		{10000000, true, false, "¥10,000,000"},
		{100000000, true, false, "¥100,000,000"},
		{1000000000, true, false, "¥1,000,000,000"},
		{10000000000, true, false, "¥10,000,000,000"},
		{100000000000, true, false, "¥100,000,000,000"},
		{1000000000000, true, false, "¥1,000,000,000,000"},
		{10000000000000, true, false, "¥10,000,000,000,000"},
		{99999999999999, true, false, "¥99,999,999,999,999"},
		{100000000000000, true, false, "¥0"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{1000, false, false, "1,000"},
		{10000, false, false, "10,000"},
		{100000, false, false, "100,000"},
		{1000000, false, false, "1,000,000"},
		{10000000, false, false, "10,000,000"},
		{100000000, false, false, "100,000,000"},
		{1000000000, false, false, "1,000,000,000"},
		{10000000000, false, false, "10,000,000,000"},
		{100000000000, false, false, "100,000,000,000"},
		{1000000000000, false, false, "1,000,000,000,000"},
		{10000000000000, false, false, "10,000,000,000,000"},
		{99999999999999, false, false, "99,999,999,999,999"},
		{100000000000000, false, false, "0"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXUIntptr(t *testing.T) {
	tests := []struct {
		price          uintptr
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1000, true, true, "¥1,000-"},
		{10000, true, true, "¥10,000-"},
		{100000, true, true, "¥100,000-"},
		{1000000, true, true, "¥1,000,000-"},
		{10000000, true, true, "¥10,000,000-"},
		{100000000, true, true, "¥100,000,000-"},
		{1000000000, true, true, "¥1,000,000,000-"},
		{10000000000, true, true, "¥10,000,000,000-"},
		{100000000000, true, true, "¥100,000,000,000-"},
		{1000000000000, true, true, "¥1,000,000,000,000-"},
		{10000000000000, true, true, "¥10,000,000,000,000-"},
		{99999999999999, true, true, "¥99,999,999,999,999-"},
		{100000000000000, true, true, "¥0-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1000, false, true, "1,000-"},
		{10000, false, true, "10,000-"},
		{100000, false, true, "100,000-"},
		{1000000, false, true, "1,000,000-"},
		{10000000, false, true, "10,000,000-"},
		{100000000, false, true, "100,000,000-"},
		{1000000000, false, true, "1,000,000,000-"},
		{10000000000, false, true, "10,000,000,000-"},
		{100000000000, false, true, "100,000,000,000-"},
		{1000000000000, false, true, "1,000,000,000,000-"},
		{10000000000000, false, true, "10,000,000,000,000-"},
		{99999999999999, false, true, "99,999,999,999,999-"},
		{100000000000000, false, true, "0-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1000, true, false, "¥1,000"},
		{10000, true, false, "¥10,000"},
		{100000, true, false, "¥100,000"},
		{1000000, true, false, "¥1,000,000"},
		{10000000, true, false, "¥10,000,000"},
		{100000000, true, false, "¥100,000,000"},
		{1000000000, true, false, "¥1,000,000,000"},
		{10000000000, true, false, "¥10,000,000,000"},
		{100000000000, true, false, "¥100,000,000,000"},
		{1000000000000, true, false, "¥1,000,000,000,000"},
		{10000000000000, true, false, "¥10,000,000,000,000"},
		{99999999999999, true, false, "¥99,999,999,999,999"},
		{100000000000000, true, false, "¥0"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{1000, false, false, "1,000"},
		{10000, false, false, "10,000"},
		{100000, false, false, "100,000"},
		{1000000, false, false, "1,000,000"},
		{10000000, false, false, "10,000,000"},
		{100000000, false, false, "100,000,000"},
		{1000000000, false, false, "1,000,000,000"},
		{10000000000, false, false, "10,000,000,000"},
		{100000000000, false, false, "100,000,000,000"},
		{1000000000000, false, false, "1,000,000,000,000"},
		{10000000000000, false, false, "10,000,000,000,000"},
		{99999999999999, false, false, "99,999,999,999,999"},
		{100000000000000, false, false, "0"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXArbitraryType(t *testing.T) {
	type MyInt int
	tests := []struct {
		price          MyInt
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1, true, true, "¥1-"},
		{10, true, true, "¥10-"},
		{100, true, true, "¥100-"},
		{1000, true, true, "¥1,000-"},
		{10000, true, true, "¥10,000-"},
		{100000, true, true, "¥100,000-"},
		{1000000, true, true, "¥1,000,000-"},
		{10000000, true, true, "¥10,000,000-"},
		{100000000, true, true, "¥100,000,000-"},
		{1000000000, true, true, "¥1,000,000,000-"},
		{10000000000, true, true, "¥10,000,000,000-"},
		{100000000000, true, true, "¥100,000,000,000-"},
		{1000000000000, true, true, "¥1,000,000,000,000-"},
		{10000000000000, true, true, "¥10,000,000,000,000-"},
		{99999999999999, true, true, "¥99,999,999,999,999-"},
		{100000000000000, true, true, "¥0-"},
		{-1, true, true, "¥-1-"},
		{-10, true, true, "¥-10-"},
		{-100, true, true, "¥-100-"},
		{-1000, true, true, "¥-1,000-"},
		{-10000, true, true, "¥-10,000-"},
		{-100000, true, true, "¥-100,000-"},
		{-1000000, true, true, "¥-1,000,000-"},
		{-10000000, true, true, "¥-10,000,000-"},
		{-100000000, true, true, "¥-100,000,000-"},
		{-1000000000, true, true, "¥-1,000,000,000-"},
		{-10000000000, true, true, "¥-10,000,000,000-"},
		{-100000000000, true, true, "¥-100,000,000,000-"},
		{-1000000000000, true, true, "¥-1,000,000,000,000-"},
		{-10000000000000, true, true, "¥-10,000,000,000,000-"},
		{-99999999999999, true, true, "¥-99,999,999,999,999-"},
		{-100000000000000, true, true, "¥0-"},
		{1, false, true, "1-"},
		{10, false, true, "10-"},
		{100, false, true, "100-"},
		{1000, false, true, "1,000-"},
		{10000, false, true, "10,000-"},
		{100000, false, true, "100,000-"},
		{1000000, false, true, "1,000,000-"},
		{10000000, false, true, "10,000,000-"},
		{100000000, false, true, "100,000,000-"},
		{1000000000, false, true, "1,000,000,000-"},
		{10000000000, false, true, "10,000,000,000-"},
		{100000000000, false, true, "100,000,000,000-"},
		{1000000000000, false, true, "1,000,000,000,000-"},
		{10000000000000, false, true, "10,000,000,000,000-"},
		{99999999999999, false, true, "99,999,999,999,999-"},
		{100000000000000, false, true, "0-"},
		{-1, false, true, "-1-"},
		{-10, false, true, "-10-"},
		{-100, false, true, "-100-"},
		{-1000, false, true, "-1,000-"},
		{-10000, false, true, "-10,000-"},
		{-100000, false, true, "-100,000-"},
		{-1000000, false, true, "-1,000,000-"},
		{-10000000, false, true, "-10,000,000-"},
		{-100000000, false, true, "-100,000,000-"},
		{-1000000000, false, true, "-1,000,000,000-"},
		{-10000000000, false, true, "-10,000,000,000-"},
		{-100000000000, false, true, "-100,000,000,000-"},
		{-1000000000000, false, true, "-1,000,000,000,000-"},
		{-10000000000000, false, true, "-10,000,000,000,000-"},
		{-99999999999999, false, true, "-99,999,999,999,999-"},
		{-100000000000000, false, true, "0-"},
		{1, true, false, "¥1"},
		{10, true, false, "¥10"},
		{100, true, false, "¥100"},
		{1000, true, false, "¥1,000"},
		{10000, true, false, "¥10,000"},
		{100000, true, false, "¥100,000"},
		{1000000, true, false, "¥1,000,000"},
		{10000000, true, false, "¥10,000,000"},
		{100000000, true, false, "¥100,000,000"},
		{1000000000, true, false, "¥1,000,000,000"},
		{10000000000, true, false, "¥10,000,000,000"},
		{100000000000, true, false, "¥100,000,000,000"},
		{1000000000000, true, false, "¥1,000,000,000,000"},
		{10000000000000, true, false, "¥10,000,000,000,000"},
		{99999999999999, true, false, "¥99,999,999,999,999"},
		{100000000000000, true, false, "¥0"},
		{-1, true, false, "¥-1"},
		{-10, true, false, "¥-10"},
		{-100, true, false, "¥-100"},
		{-1000, true, false, "¥-1,000"},
		{-10000, true, false, "¥-10,000"},
		{-100000, true, false, "¥-100,000"},
		{-1000000, true, false, "¥-1,000,000"},
		{-10000000, true, false, "¥-10,000,000"},
		{-100000000, true, false, "¥-100,000,000"},
		{-1000000000, true, false, "¥-1,000,000,000"},
		{-10000000000, true, false, "¥-10,000,000,000"},
		{-100000000000, true, false, "¥-100,000,000,000"},
		{-1000000000000, true, false, "¥-1,000,000,000,000"},
		{-10000000000000, true, false, "¥-10,000,000,000,000"},
		{-99999999999999, true, false, "¥-99,999,999,999,999"},
		{-100000000000000, true, false, "¥0"},
		{1, false, false, "1"},
		{10, false, false, "10"},
		{100, false, false, "100"},
		{1000, false, false, "1,000"},
		{10000, false, false, "10,000"},
		{100000, false, false, "100,000"},
		{1000000, false, false, "1,000,000"},
		{10000000, false, false, "10,000,000"},
		{100000000, false, false, "100,000,000"},
		{1000000000, false, false, "1,000,000,000"},
		{10000000000, false, false, "10,000,000,000"},
		{100000000000, false, false, "100,000,000,000"},
		{1000000000000, false, false, "1,000,000,000,000"},
		{10000000000000, false, false, "10,000,000,000,000"},
		{99999999999999, false, false, "99,999,999,999,999"},
		{100000000000000, false, false, "0"},
		{-1, false, false, "-1"},
		{-10, false, false, "-10"},
		{-100, false, false, "-100"},
		{-1000, false, false, "-1,000"},
		{-10000, false, false, "-10,000"},
		{-100000, false, false, "-100,000"},
		{-1000000, false, false, "-1,000,000"},
		{-10000000, false, false, "-10,000,000"},
		{-100000000, false, false, "-100,000,000"},
		{-1000000000, false, false, "-1,000,000,000"},
		{-10000000000, false, false, "-10,000,000,000"},
		{-100000000000, false, false, "-100,000,000,000"},
		{-1000000000000, false, false, "-1,000,000,000,000"},
		{-10000000000000, false, false, "-10,000,000,000,000"},
		{-99999999999999, false, false, "-99,999,999,999,999"},
		{-100000000000000, false, false, "0"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %d, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}

func TestFormatXFloat64(t *testing.T) {
	tests := []struct {
		price          float64
		prefixEnabled  bool
		suffixEnabled  bool
		expectedResult string
	}{
		{1.0, true, true, "¥1-"},
		{10.0, true, true, "¥10-"},
		{100.0, true, true, "¥100-"},
		{1000.0, true, true, "¥1,000-"},
		{10000.0, true, true, "¥10,000-"},
		{100000.0, true, true, "¥100,000-"},
		{1000000.0, true, true, "¥1,000,000-"},
		{10000000.0, true, true, "¥10,000,000-"},
		{100000000.0, true, true, "¥100,000,000-"},
		{1000000000.0, true, true, "¥1,000,000,000-"},
		{10000000000.0, true, true, "¥10,000,000,000-"},
		{100000000000.0, true, true, "¥100,000,000,000-"},
		{1000000000000.0, true, true, "¥1,000,000,000,000-"},
		{10000000000000.0, true, true, "¥10,000,000,000,000-"},
		{99999999999999.0, true, true, "¥99,999,999,999,999-"},
		{100000000000000.0, true, true, "¥0-"},
		{-1.0, true, true, "¥-1-"},
		{-10.0, true, true, "¥-10-"},
		{-100.0, true, true, "¥-100-"},
		{-1000.0, true, true, "¥-1,000-"},
		{-10000.0, true, true, "¥-10,000-"},
		{-100000.0, true, true, "¥-100,000-"},
		{-1000000.0, true, true, "¥-1,000,000-"},
		{-10000000.0, true, true, "¥-10,000,000-"},
		{-100000000.0, true, true, "¥-100,000,000-"},
		{-1000000000.0, true, true, "¥-1,000,000,000-"},
		{-10000000000.0, true, true, "¥-10,000,000,000-"},
		{-100000000000.0, true, true, "¥-100,000,000,000-"},
		{-1000000000000.0, true, true, "¥-1,000,000,000,000-"},
		{-10000000000000.0, true, true, "¥-10,000,000,000,000-"},
		{-99999999999999.0, true, true, "¥-99,999,999,999,999-"},
		{-100000000000000.0, true, true, "¥0-"},
		{1.0, false, true, "1-"},
		{10.0, false, true, "10-"},
		{100.0, false, true, "100-"},
		{1000.0, false, true, "1,000-"},
		{10000.0, false, true, "10,000-"},
		{100000.0, false, true, "100,000-"},
		{1000000.0, false, true, "1,000,000-"},
		{10000000.0, false, true, "10,000,000-"},
		{100000000.0, false, true, "100,000,000-"},
		{1000000000.0, false, true, "1,000,000,000-"},
		{10000000000.0, false, true, "10,000,000,000-"},
		{100000000000.0, false, true, "100,000,000,000-"},
		{1000000000000.0, false, true, "1,000,000,000,000-"},
		{10000000000000.0, false, true, "10,000,000,000,000-"},
		{99999999999999.0, false, true, "99,999,999,999,999-"},
		{100000000000000.0, false, true, "0-"},
		{-1.0, false, true, "-1-"},
		{-10.0, false, true, "-10-"},
		{-100.0, false, true, "-100-"},
		{-1000.0, false, true, "-1,000-"},
		{-10000.0, false, true, "-10,000-"},
		{-100000.0, false, true, "-100,000-"},
		{-1000000.0, false, true, "-1,000,000-"},
		{-10000000.0, false, true, "-10,000,000-"},
		{-100000000.0, false, true, "-100,000,000-"},
		{-1000000000.0, false, true, "-1,000,000,000-"},
		{-10000000000.0, false, true, "-10,000,000,000-"},
		{-100000000000.0, false, true, "-100,000,000,000-"},
		{-1000000000000.0, false, true, "-1,000,000,000,000-"},
		{-10000000000000.0, false, true, "-10,000,000,000,000-"},
		{-99999999999999.0, false, true, "-99,999,999,999,999-"},
		{-100000000000000.0, false, true, "0-"},
		{1.0, true, false, "¥1"},
		{10.0, true, false, "¥10"},
		{100.0, true, false, "¥100"},
		{1000.0, true, false, "¥1,000"},
		{10000.0, true, false, "¥10,000"},
		{100000.0, true, false, "¥100,000"},
		{1000000.0, true, false, "¥1,000,000"},
		{10000000.0, true, false, "¥10,000,000"},
		{100000000.0, true, false, "¥100,000,000"},
		{1000000000.0, true, false, "¥1,000,000,000"},
		{10000000000.0, true, false, "¥10,000,000,000"},
		{100000000000.0, true, false, "¥100,000,000,000"},
		{1000000000000.0, true, false, "¥1,000,000,000,000"},
		{10000000000000.0, true, false, "¥10,000,000,000,000"},
		{99999999999999.0, true, false, "¥99,999,999,999,999"},
		{100000000000000.0, true, false, "¥0"},
		{-1.0, true, false, "¥-1"},
		{-10.0, true, false, "¥-10"},
		{-100.0, true, false, "¥-100"},
		{-1000.0, true, false, "¥-1,000"},
		{-10000.0, true, false, "¥-10,000"},
		{-100000.0, true, false, "¥-100,000"},
		{-1000000.0, true, false, "¥-1,000,000"},
		{-10000000.0, true, false, "¥-10,000,000"},
		{-100000000.0, true, false, "¥-100,000,000"},
		{-1000000000.0, true, false, "¥-1,000,000,000"},
		{-10000000000.0, true, false, "¥-10,000,000,000"},
		{-100000000000.0, true, false, "¥-100,000,000,000"},
		{-1000000000000.0, true, false, "¥-1,000,000,000,000"},
		{-10000000000000.0, true, false, "¥-10,000,000,000,000"},
		{-99999999999999.0, true, false, "¥-99,999,999,999,999"},
		{-100000000000000.0, true, false, "¥0"},
		{1.0, false, false, "1"},
		{10.0, false, false, "10"},
		{100.0, false, false, "100"},
		{1000.0, false, false, "1,000"},
		{10000.0, false, false, "10,000"},
		{100000.0, false, false, "100,000"},
		{1000000.0, false, false, "1,000,000"},
		{10000000.0, false, false, "10,000,000"},
		{100000000.0, false, false, "100,000,000"},
		{1000000000.0, false, false, "1,000,000,000"},
		{10000000000.0, false, false, "10,000,000,000"},
		{100000000000.0, false, false, "100,000,000,000"},
		{1000000000000.0, false, false, "1,000,000,000,000"},
		{10000000000000.0, false, false, "10,000,000,000,000"},
		{99999999999999.0, false, false, "99,999,999,999,999"},
		{100000000000000.0, false, false, "0"},
		{-1.0, false, false, "-1"},
		{-10.0, false, false, "-10"},
		{-100.0, false, false, "-100"},
		{-1000.0, false, false, "-1,000"},
		{-10000.0, false, false, "-10,000"},
		{-100000.0, false, false, "-100,000"},
		{-1000000.0, false, false, "-1,000,000"},
		{-10000000.0, false, false, "-10,000,000"},
		{-100000000.0, false, false, "-100,000,000"},
		{-1000000000.0, false, false, "-1,000,000,000"},
		{-10000000000.0, false, false, "-10,000,000,000"},
		{-100000000000.0, false, false, "-100,000,000,000"},
		{-1000000000000.0, false, false, "-1,000,000,000,000"},
		{-10000000000000.0, false, false, "-10,000,000,000,000"},
		{-99999999999999.0, false, false, "-99,999,999,999,999"},
		{-100000000000000.0, false, false, "0"},
	}

	for _, test := range tests {
		result := formatter.FormatX(test.price, test.prefixEnabled, test.suffixEnabled)
		if result != test.expectedResult {
			t.Errorf("Test failed. Price: %f, Expected resutlt : %s, Actucal result %s", test.price, test.expectedResult, result)
		} else {
			fmt.Println("Test passed.")
		}
	}
}
