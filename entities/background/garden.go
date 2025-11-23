package background

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type GardenCreator struct {
	loader          resources.IResourceLoader
	TargetImageSize gmath.Vec
}

func NewGardenCreator(loader resources.IResourceLoader) *GardenCreator {
	return &GardenCreator{
		loader:          loader,
		TargetImageSize: gmath.Vec{X: 20, Y: 20},
	}
}

func (c GardenCreator) Create(world donburi.World, position components.PositionData) (donburi.Entity, error) {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.RectCollider,
		components.Garden,
	)

	en := world.Entry(entity)

	im, err := c.loader.LoadImage(resources.ImageGarden)
	if err != nil {
		return 0, err
	}

	components.Sprite.SetValue(en, components.SpriteData{
		Image: images.Image{
			Image: im, Scale: render.GetImageScale(im.Bounds(), c.TargetImageSize),
		},
		Z: 1,
	})

	rect := gmath.Rect{
		Min: c.TargetImageSize.Mulf(-0.5),
		Max: c.TargetImageSize.Mulf(0.5),
	}
	components.RectCollider.SetValue(en, components.RectColliderData{Rect: rect})
	components.Position.SetValue(en, position)

	return entity, nil
}
