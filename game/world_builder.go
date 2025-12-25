package game

import (
	"image"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/utils"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

var (
	WorldSize = image.Rect(0, 0, 600, 600)
)

type WorldBuilder struct {
	grassCreator background.GrassCreator
	coalCreator  entities.CoalCreator
}

func NewWorldBuilder(
	grassCreator background.GrassCreator,
	coalCreator entities.CoalCreator,
) *WorldBuilder {
	return &WorldBuilder{
		grassCreator: grassCreator,
		coalCreator:  coalCreator,
	}
}

func (b WorldBuilder) Build(world donburi.World) error {
	currentY := WorldSize.Min.Y
	for currentY <= WorldSize.Max.Y {
		currentX := WorldSize.Min.X
		for currentX <= WorldSize.Max.X {
			b.grassCreator.Create(world, components.PositionData{
				Vec: gmath.Vec{X: float64(currentX), Y: float64(currentY)},
			})
			currentX += int(b.grassCreator.TargetImageSize.X)
		}
		currentY += int(b.grassCreator.TargetImageSize.Y)
	}

	for range 15 {
		b.coalCreator.Create(world, components.PositionData{
			Vec: gmath.Vec{X: utils.RandomFloat(0, WorldSize.Max.X), Y: utils.RandomFloat(0, WorldSize.Max.Y)},
		})
	}

	return nil
}
