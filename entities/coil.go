package entities

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/constants/z"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type CoilCreator struct {
	loader          resources.IResourceLoader
	TargetImageSize gmath.Vec
}

func NewCoilCreator(loader resources.IResourceLoader) *CoilCreator {
	return &CoilCreator{
		loader:          loader,
		TargetImageSize: gmath.Vec{X: 25, Y: 25},
	}
}

func (c CoilCreator) Create(world donburi.World, position components.PositionData) donburi.Entity {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.SpriteCollider,
		components.Obstacle,
		components.Destroyable,
	)

	en := world.Entry(entity)

	im := c.loader.LoadImage(resources.ImageCoil)

	sprite := components.SpriteData{
		Image: images.Image{
			Image: im,
			Scale: render.GetImageScale(im.Bounds(), c.TargetImageSize),
		},
		Z: z.OBJ,
	}
	components.Sprite.SetValue(en, sprite)

	components.Position.SetValue(en, position)

	rect := utils.GetRectOfBottomOfParent(c.TargetImageSize, 0.3)

	components.SpriteCollider.SetValue(en, components.SpriteColliderData{
		ActiveZone: &rect,
	})

	return entity
}
