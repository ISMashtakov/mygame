package actions

import (
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type GardenCreatingRequestData struct {
	Point gmath.Vec
}

var GardenCreatingRequest = donburi.NewComponentType[GardenCreatingRequestData]()

type PickaxeHitRequestData struct {
	Point gmath.Vec
}

var PickaxeHitRequest = donburi.NewComponentType[PickaxeHitRequestData]()
