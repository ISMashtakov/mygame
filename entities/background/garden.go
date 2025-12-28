package background

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/constants/z"
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

func (c GardenCreator) Create(world donburi.World, position components.PositionData) *donburi.Entry {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.RectCollider,
		components.Garden,
	)

	en := world.Entry(entity)

	im := c.loader.LoadImage(resources.ImageGarden)

	components.Sprite.SetValue(en, components.SpriteData{
		Image: images.Image{
			Image: im, Scale: render.GetImageScale(im.Bounds(), c.TargetImageSize),
		},
		Z: z.GARDEN,
	})

	rect := gmath.Rect{
		Min: c.TargetImageSize.Mulf(-0.5),
		Max: c.TargetImageSize.Mulf(0.5),
	}
	components.RectCollider.SetValue(en, components.RectColliderData{Rect: rect})
	components.Position.SetValue(en, position)

	return en
}
