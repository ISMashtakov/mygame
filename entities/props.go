package entities

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/constants/z"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type PropsCreator struct {
	TargetImageSize  gmath.Vec
	RectColliderSize gmath.Vec
}

func NewPropsCreator() *PropsCreator {
	return &PropsCreator{
		TargetImageSize:  gmath.Vec{X: 25, Y: 25},
		RectColliderSize: gmath.Vec{X: 10, Y: 10},
	}
}

func (c PropsCreator) Create(world donburi.World, item items.IItem, pos gmath.Vec) *donburi.Entry {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.RectCollider,
		components.Prop,
		components.DisabledColliders,
	)

	en := world.Entry(entity)

	components.Sprite.SetValue(en, components.SpriteData{
		Image: images.Image{
			Image: item.GetImage(),
			Scale: render.GetImageScale(item.GetImage().Bounds(), c.TargetImageSize),
		},
		Z: z.PROP,
	})

	components.Position.SetValue(en, components.PositionData{Vec: pos})

	components.RectCollider.SetValue(en, components.RectColliderData{
		Rect: gmath.Rect{Min: c.RectColliderSize.Mulf(-0.5), Max: c.RectColliderSize.Mulf(0.5)},
	})

	components.Prop.Set(en, &components.PropData{Item: item})

	return en
}
