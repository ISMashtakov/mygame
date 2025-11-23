package entities

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type StoneCreator struct {
	loader          resources.IResourceLoader
	TargetImageSize gmath.Vec
}

func NewStoneCreator(loader resources.IResourceLoader) *StoneCreator {
	return &StoneCreator{
		loader:          loader,
		TargetImageSize: gmath.Vec{X: 25, Y: 25},
	}
}

func (c StoneCreator) Create(world donburi.World, position components.PositionData) (donburi.Entity, error) {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.SpriteCollider,
		components.Obstacle,
	)

	en := world.Entry(entity)

	im, err := c.loader.LoadImage(resources.ImageStone)
	if err != nil {
		return 0, err
	}

	sprite := components.SpriteData{
		Image: images.Image{
			Image: im,
			Scale: render.GetImageScale(im.Bounds(), c.TargetImageSize),
		},
		Z: 5,
	}
	components.Sprite.SetValue(en, sprite)

	components.Position.SetValue(en, position)

	return entity, nil
}
