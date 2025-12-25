package entities

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/constants/z"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type CoalCreator struct {
	loader          resources.IResourceLoader
	TargetImageSize gmath.Vec
	itemsFactory    items.Factory
}

func NewCoalCreator(loader resources.IResourceLoader, itemsFactory items.Factory) *CoalCreator {
	return &CoalCreator{
		loader:          loader,
		TargetImageSize: gmath.Vec{X: 25, Y: 25},
		itemsFactory:    itemsFactory,
	}
}

func (c CoalCreator) Create(world donburi.World, position components.PositionData) donburi.Entity {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.SpriteCollider,
		components.Obstacle,
		components.Destroyable,
	)

	en := world.Entry(entity)

	im := c.loader.LoadImage(resources.ImageCoal)

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

	components.Destroyable.SetValue(en, components.DestroyableData{
		Resources: utils.SlicsByFunc(utils.RandomInt(1, 4), c.itemsFactory.Coal),
	})

	return entity
}
