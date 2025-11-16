package entities

import (
	"image"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

const (
	sizeX float64 = 17
	sizeY float64 = 25
)

func CreateCharacter(world donburi.World, resourceLoader resources.IResourceLoader) (*donburi.Entry, error) {
	entity := world.Create(components.Position, components.Sprite)
	entry := world.Entry(entity)

	resImage, err := resourceLoader.LoadImage(resources.ImageCharacter)
	if err != nil {
		return nil, err
	}

	rect := image.Rect(
		80, 60, 250, 310,
	)

	subImage := resImage.SubImage(rect)

	components.Sprite.SetValue(entry, components.SpriteData{
		Image: subImage.(*ebiten.Image),
		Scale: &gmath.Vec{X: sizeX / float64(rect.Dx()), Y: sizeY / float64(rect.Dy())},
	})

	return entry, nil
}
