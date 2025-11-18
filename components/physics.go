package components

import (
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type SpeedData struct {
	gmath.Vec
}

var Speed = donburi.NewComponentType[SpeedData]()
