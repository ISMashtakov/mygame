package game

import (
	"fmt"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Game struct {
	systems []core.ISystem
	world   donburi.World
}

func NewGame(world donburi.World, systems []core.ISystem) *Game {
	return &Game{
		world:   world,
		systems: systems,
	}

}

func (g *Game) Update() error {
	for _, system := range g.systems {
		if err := system.Update(g.world); err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for en := range donburi.NewQuery(filter.Contains(components.Sprite, components.Position)).Iter(g.world) {
		sprite, position := components.Sprite.Get(en), components.Position.Get(en)

		op := ebiten.DrawImageOptions{}

		if !sprite.Scale.IsZero() {
			op.GeoM.Scale(sprite.Scale.X, sprite.Scale.Y)
		}
		op.GeoM.Translate(position.X, position.Y)

		screen.DrawImage(sprite.Image, &op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
