package entities

import (
	"github.com/ISMashtakov/mygame/components"
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

	return entity, nil
}
