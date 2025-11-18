package character

import (
	"image"

	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/physics"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
	"github.com/quasilyte/gmath"
)

const (
	sizeX float64 = 17
	sizeY float64 = 25
)

type CharacterCreator struct {
	mapper         *ecs.Map4[core.Position, core.Sprite, physics.Speed, Character]
	resourceLoader resources.IResourceLoader
}

func NewCharacterCreator(world *ecs.World, resourceLoader resources.IResourceLoader) *CharacterCreator {
	return &CharacterCreator{
		mapper:         ecs.NewMap4[core.Position, core.Sprite, physics.Speed, Character](world),
		resourceLoader: resourceLoader,
	}
}

func (c CharacterCreator) Create() (ecs.Entity, error) {
	resImage, err := c.resourceLoader.LoadImage(resources.ImageCharacter)
	if err != nil {
		return ecs.Entity{}, err
	}

	rect := image.Rect(
		80, 60, 250, 310,
	)

	subImage := resImage.SubImage(rect)

	return c.mapper.NewEntity(
		&core.Position{},
		&core.Sprite{
			Image: subImage.(*ebiten.Image),
			Scale: &gmath.Vec{X: sizeX / float64(rect.Dx()), Y: sizeY / float64(rect.Dy())},
		},
		&physics.Speed{},
		&Character{},
	), nil
}
