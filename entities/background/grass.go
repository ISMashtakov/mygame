package background

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/constants"
	"github.com/ISMashtakov/mygame/constants/z"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type GrassCreator struct {
	loader          resources.IResourceLoader
	TargetImageSize gmath.Vec
}

func NewGrassCreator(loader resources.IResourceLoader) *GrassCreator {
	return &GrassCreator{
		loader:          loader,
		TargetImageSize: constants.TileSize,
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

	components.Sprite.SetValue(en, components.SpriteData{
		Image: images.Image{
			Image: im, Scale: render.GetImageScale(im.Bounds(), c.TargetImageSize),
		},
		Z: z.GRASS,
	})

	components.Position.SetValue(en, position)

	return entity, nil
}
