package utils

import "math/rand/v2"

func RandomInt(from, to int) int {
	return from + (rand.Int() % (to - from))
}

func RandomFloat(from, to int) float64 {
	return float64(RandomInt(from, to-1)) + rand.Float64()
}
