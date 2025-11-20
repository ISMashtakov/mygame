package components

import (
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type MovementRequestData struct {
	gmath.Vec
}

var MovementRequest = donburi.NewComponentType[MovementRequestData]()
