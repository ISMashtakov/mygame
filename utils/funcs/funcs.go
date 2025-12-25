package funcs

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Func func(x float64) float64

func Zero() Func {
	return func(x float64) float64 {
		return 0
	}
}

func Line(a float64) Func {
	return func(x float64) float64 {
		return a * x
	}
}

func LineTo(x, y float64) Func {
	return Line(y / x)
}

func Square(a, b float64) Func {
	return func(x float64) float64 {
		q := a*x*x + b*x
		return q
	}
}

func SquareTo(x1, y1, x2, y2 float64) Func {
	a := (x2*y1 - x1*y2) / (x1 - x2) / (x1 * x2)
	b := (y2-y1)/(x2-x1) - a*(x1+x2)

	return Square(a, b)
}

func Abs(x1, y1 float64) Func {
	// k(x -x1) + y
	return func(x float64) float64 {
		return (-y1/x1)*math.Abs(x-x1) + y1
	}
}

func X(duration time.Duration) float64 {
	return float64(ebiten.TPS()) * duration.Seconds()
}
