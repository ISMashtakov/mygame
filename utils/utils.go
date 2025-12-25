package utils

import (
	"github.com/quasilyte/gmath"
	"github.com/samber/lo"
)

func SlicsByFunc[T any](count int, fun func() T) []T {
	var res []T
	for range count {
		res = append(res, fun())
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

func GetRectOfBottomOfParent(vec gmath.Vec, percent float64) gmath.Rect {
	offset := vec.
		Mulf(-0.5).
		Add(gmath.Vec{Y: vec.Y * percent}).
		Add(gmath.Vec{Y: vec.Y * 0.25}) // TODO: Я не понимаю, почему нужен этот сдвиг

	return gmath.Rect{
		Min: offset,
		Max: vec.Sub(gmath.Vec{Y: vec.Y * 0.5}).Add(offset),
	}
}
