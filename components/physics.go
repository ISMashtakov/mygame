package components

import (
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type MovementRequestData struct {
	gmath.Vec
}

var MovementRequest = donburi.NewComponentType[MovementRequestData]()

type RectColliderData struct {
	gmath.Rect
}

var RectCollider = donburi.NewComponentType[RectColliderData]()

type SpriteColliderData struct {
	ActiveZone *gmath.Rect
}

var SpriteCollider = donburi.NewComponentType[SpriteColliderData]()
