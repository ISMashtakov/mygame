package game

import (
	"github.com/ISMashtakov/mygame/constants"
	"github.com/ISMashtakov/mygame/gui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Game struct {
	systems  []ISystem
	world    donburi.World
	renderer Renderer
	gui      *gui.GUI
}

func NewGame(renderer Renderer, world donburi.World, systems []ISystem, gui *gui.GUI) *Game {
	return &Game{
		world:    world,
		systems:  systems,
		renderer: renderer,
		gui:      gui,
	}
}

func (g *Game) Update() error {
	for _, system := range g.systems {
		system.Update(g.world)
	}

	g.gui.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Draw(screen, g.world)
	g.gui.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return int(constants.TargetLayout.X), int(constants.TargetLayout.Y)
}
