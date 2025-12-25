package components

import (
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type MovementData struct {
	gmath.Vec
}

var Movement = donburi.NewComponentType[MovementData]()

type RectColliderData struct {
	gmath.Rect
}

var RectCollider = donburi.NewComponentType[RectColliderData]()

type SpriteColliderData struct {
	ActiveZone *gmath.Rect
}

var SpriteCollider = donburi.NewComponentType[SpriteColliderData]()

var DisabledColliders = donburi.NewTag()
