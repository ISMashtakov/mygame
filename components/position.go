package components

import (
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type PositionData struct {
	gmath.Vec
}

var Position = donburi.NewComponentType[PositionData]()
