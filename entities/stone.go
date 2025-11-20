package entities

import (
	"image"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/yohamta/donburi"
)

type StoneCreator struct {
	loader          resources.IResourceLoader
	TargetImageSize image.Rectangle
}

func NewStoneCreator(loader resources.IResourceLoader) *StoneCreator {
	return &StoneCreator{
		loader:          loader,
		TargetImageSize: image.Rect(0, 0, 25, 25),
	}
}

func (c StoneCreator) Create(world donburi.World, position components.PositionData) (donburi.Entity, error) {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.Collider,
	)

	en := world.Entry(entity)

	im, err := c.loader.LoadImage(resources.ImageStone)
	if err != nil {
		return 0, err
	}

	sprite := components.SpriteData{
		Image: im,
		Scale: render.GetImageScale(im.Bounds(), c.TargetImageSize),
	}
	components.Sprite.SetValue(en, sprite)

	components.Position.SetValue(en, position)
	components.Collider.SetValue(en, *components.GetColliderDataBySprite(sprite))

	return entity, nil
}
