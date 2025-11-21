package utils

import (
	"iter"
)

func GetListByIterator[T any](iterator iter.Seq[T]) []T {
	var res []T
	for i := range iterator {
		res = append(res, i)
	}

	return res
}
