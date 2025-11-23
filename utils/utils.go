package utils

import (
	"iter"

	"github.com/quasilyte/gmath"
	"github.com/samber/lo"
)

func GetListByIterator[T any](iterator iter.Seq[T]) []T {
	var res []T
	for i := range iterator {
		res = append(res, i)
	}

	return res
}

func FloorByNearestStep(num, step int) int {
	fail := num % step

	return lo.Ternary(fail <= step/2, num-fail, num+(step-fail))
}

func FloorByNearestStepVec(vec, stepVec gmath.Vec) gmath.Vec {
	return gmath.Vec{
		X: float64(FloorByNearestStep(int(vec.X), int(stepVec.X))),
		Y: float64(FloorByNearestStep(int(vec.Y), int(stepVec.Y))),
	}
}
