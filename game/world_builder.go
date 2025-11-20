package game

import (
	"image"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

var (
	WorldSize = image.Rect(0, 0, 500, 500)
)

type WorldBuilder struct {
	grassCreator background.GrassCreator
}

func NewWorldBuilder(grassCreator background.GrassCreator) *WorldBuilder {
	return &WorldBuilder{
		grassCreator: grassCreator,
	}
}

func (b WorldBuilder) Build(world donburi.World) error {
	currentY := WorldSize.Min.Y
	for currentY <= WorldSize.Max.Y {
		currentX := WorldSize.Min.X
		for currentX <= WorldSize.Max.X {
			_, err := b.grassCreator.Create(world, components.PositionData{
				Vec: gmath.Vec{X: float64(currentX), Y: float64(currentY)},
			})
			if err != nil {
				return err
			}
			currentX += b.grassCreator.TargetImageSize.Dx()
		}
		currentY += b.grassCreator.TargetImageSize.Dy()
	}

	return nil
}
