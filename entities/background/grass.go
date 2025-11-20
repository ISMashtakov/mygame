package background

import (
	"image"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/yohamta/donburi"
)

var ()

type GrassCreator struct {
	loader          resources.IResourceLoader
	TargetImageSize image.Rectangle
}

func NewGrassCreator(loader resources.IResourceLoader) *GrassCreator {
	return &GrassCreator{
		loader:          loader,
		TargetImageSize: image.Rect(0, 0, 25, 25),
	}
}

func (c GrassCreator) Create(world donburi.World, position components.PositionData) (donburi.Entity, error) {
	entity := world.Create(
		components.Position,
		components.Sprite,
	)

	en := world.Entry(entity)

	im, err := c.loader.LoadImage(resources.ImageGrass)
	if err != nil {
		return 0, err
	}

	components.Sprite.Set(en, &components.SpriteData{
		Image: im,
		Scale: render.GetImageScale(im.Bounds(), c.TargetImageSize),
	})

	components.Position.Set(en, &position)

	return entity, nil
}
