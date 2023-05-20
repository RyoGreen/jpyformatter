package formatter

import (
	"errors"
	"strconv"

	"golang.org/x/exp/constraints"
)

var ErrPriceExceedsLimit = errors.New("error: price exceeds 1000000000000 yen")

type unit interface {
	constraints.Signed | constraints.Float
}

func Format[U unit](price U, prefixEnabled, suffixEnabled bool) (string, error) {
	p := int(price)
	if p >= 1000000000000 {
		return "", ErrPriceExceedsLimit
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
	if prefixEnabled {
		result = append([]byte{0xC2, 0xA5}, result...)
	}
	if suffixEnabled {
		result = append(result, '-')
	}
	return string(result), nil
}
