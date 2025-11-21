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

func GetRectColliderDataBySprite(sprite SpriteData) *RectColliderData {
	rect := gmath.Rect{
		Max: gmath.VecFromStd(sprite.Image.Bounds().Max.Sub(sprite.Image.Bounds().Min)),
	}
	rect.Max = rect.Max.Mul(sprite.Scale)

	return &RectColliderData{
		Rect: rect,
	}
}

var RectCollider = donburi.NewComponentType[RectColliderData]()

type SpriteColliderData struct{}

var SpriteCollider = donburi.NewComponentType[SpriteColliderData]()
