package game

import (
	"github.com/ISMashtakov/mygame/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
)

type Game struct {
	renderFilter *ecs.Filter2[core.Sprite, core.Position]
	systems      []core.ISystem
}

func NewGame(world *ecs.World, systems []core.ISystem) *Game {
	return &Game{
		renderFilter: ecs.NewFilter2[core.Sprite, core.Position](world),
		systems:      systems,
	}

}

func (g *Game) Update() error {
	for _, system := range g.systems {
		if err := system.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	query := g.renderFilter.Query()
	for query.Next() {
		sprite, position := query.Get()

		op := ebiten.DrawImageOptions{}

		op.GeoM.Translate(position.X, position.Y)

		if sprite.Scale != nil {
			op.GeoM.Scale(sprite.Scale.X, sprite.Scale.Y)
		}

		screen.DrawImage(sprite.Image, &op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
