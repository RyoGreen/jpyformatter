package formatter

import (
	"errors"
	"math"
	"strconv"

	"golang.org/x/exp/constraints"
)

var ErrPriceExceedsLimit = errors.New("error: price must be between 99,999,999,999,999 yen and -99,999,999,999,999")

const limitPrice = 100000000000000

type unit interface {
	constraints.Signed | constraints.Float | constraints.Unsigned
}

func Format[U unit](price U, prefixEnabled, suffixEnabled bool) (string, error) {
	p := int(price)
	if p >= limitPrice {
		return "", ErrPriceExceedsLimit
	}
	var isNagative bool
	if p < 0 {
		isNagative = true
		p = int(math.Abs(float64(p)))
	}
	str := strconv.Itoa(p)
	var result []byte
	count := 0
	for i := len(str) - 1; i >= 0; i-- {
		if count%3 == 0 && len(str) != count && count != 0 {
			result = append([]byte{','}, result...)
		}
		result = append([]byte{str[i]}, result...)
		count++
	}
	if isNagative {
		result = append([]byte{45}, result...)
	}
	if prefixEnabled {
		result = append([]byte{0xC2, 0xA5}, result...)
	}
	if suffixEnabled {
		result = append(result, '-')
	}
	return string(result), nil
}
