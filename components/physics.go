package components

import (
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type MovementRequestData struct {
	gmath.Vec
}

var MovementRequest = donburi.NewComponentType[MovementRequestData]()

type ColliderData struct {
	Width  float64
	Height float64
}

var Collider = donburi.NewComponentType[ColliderData]()
