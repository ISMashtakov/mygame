package game

import (
	"github.com/ISMashtakov/mygame/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Game struct {
	systems  []core.ISystem
	world    donburi.World
	renderer Renderer
}

func NewGame(renderer Renderer, world donburi.World, systems []core.ISystem) *Game {
	return &Game{
		world:    world,
		systems:  systems,
		renderer: renderer,
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
	g.renderer.Draw(screen, g.world)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
