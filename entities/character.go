package entities

import (
	"image"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

func CreateCharacter(world donburi.World, resourceLoader resources.IResourceLoader) (*donburi.Entry, error) {
	entity := world.Create(components.Position, components.Sprite)
	entry := world.Entry(entity)

	resImage, err := resourceLoader.LoadImage(resources.ImageCharacter)
	if err != nil {
		return nil, err
	}

	subImage := resImage.SubImage(image.Rect(
		80, 60, 250, 310,
	))
	components.Sprite.SetValue(entry, components.SpriteData{
		Image: subImage.(*ebiten.Image),
	})

	return entry, nil
}
