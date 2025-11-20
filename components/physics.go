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
	gmath.Rect
}

func GetColliderDataBySprite(sprite SpriteData) *ColliderData {
	rect := gmath.Rect{
		Max: gmath.VecFromStd(sprite.Image.Bounds().Max.Sub(sprite.Image.Bounds().Min)),
	}
	rect.Max = rect.Max.Mul(sprite.Scale)

	return &ColliderData{
		Rect: rect,
	}
}

var Collider = donburi.NewComponentType[ColliderData]()
