package game

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
)

type Game struct {
	World donburi.World
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for spriteEntity := range components.Sprite.Iter(g.World) {
		sprite := components.Sprite.Get(spriteEntity)
		screen.DrawImage(sprite.Image, &ebiten.DrawImageOptions{})
	}

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
