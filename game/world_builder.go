package game

import (
	"image"
	"math/rand"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

var (
	WorldSize = image.Rect(0, 0, 600, 600)
)

type WorldBuilder struct {
	grassCreator background.GrassCreator
	stoneCreator entities.StoneCreator
}

func NewWorldBuilder(
	grassCreator background.GrassCreator,
	stoneCreator entities.StoneCreator,
) *WorldBuilder {
	return &WorldBuilder{
		grassCreator: grassCreator,
		stoneCreator: stoneCreator,
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
			currentX += int(b.grassCreator.TargetImageSize.X)
		}
		currentY += int(b.grassCreator.TargetImageSize.Y)
	}

	for i := 0; i < 15; i++ {
		_, err := b.stoneCreator.Create(world, components.PositionData{
			Vec: gmath.Vec{X: float64(rand.Intn(WorldSize.Max.X)), Y: float64(rand.Intn(WorldSize.Max.Y))},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
