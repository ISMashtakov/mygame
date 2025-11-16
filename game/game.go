package game

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/hajimehoshi/ebiten/v2"
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

		op := ebiten.DrawImageOptions{}
		if sprite.Scale != nil {
			op.GeoM.Scale(sprite.Scale.X, sprite.Scale.Y)
		}

		screen.DrawImage(sprite.Image, &op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
