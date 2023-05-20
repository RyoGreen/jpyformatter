package formatter

import (
	"golang.org/x/exp/constraints"
)

type unit interface {
	constraints.Signed | constraints.Float
}

func Format[U unit](price U, prefixEnabled, suffixEnabled bool) string {
	var result string
	return result
}
