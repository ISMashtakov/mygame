package game

import (
	"fmt"

	"github.com/ISMashtakov/mygame/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Renderer struct {
	DrawColliders bool
}

func NewRenderer() *Renderer {
	return &Renderer{}

}

func (r *Renderer) Draw(screen *ebiten.Image, world donburi.World) {
	r.drawSprites(screen, world)
	r.drawColliders(screen, world)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f, TPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

func (r *Renderer) drawSprites(screen *ebiten.Image, world donburi.World) {
	for en := range donburi.NewOrderedQuery[components.SpriteData](filter.Contains(components.Sprite)).IterOrdered(world, components.Sprite) {
		sprite := components.Sprite.Get(en)

		op := ebiten.DrawImageOptions{}

		op.GeoM.Scale(sprite.Image.Scale.X, sprite.Image.Scale.Y)

		op.GeoM.Translate(-float64(sprite.Image.Bounds().Dx())/2*sprite.Image.Scale.X, -float64(sprite.Image.Bounds().Dy())/2*sprite.Image.Scale.Y)

		if sprite.Image.Flip {
			op.GeoM.Scale(-1, 1)
		}

		if en.HasComponent(components.Position) {
			position := components.Position.Get(en)
			op.GeoM.Translate(position.X, position.Y)
		}

		screen.DrawImage(sprite.Image.Image, &op)
	}
}

func (r *Renderer) drawColliders(screen *ebiten.Image, world donburi.World) {
	if r.DrawColliders {
		for en := range donburi.NewQuery(filter.Contains(components.RectCollider)).Iter(world) {
			collider := components.RectCollider.Get(en)

			r := graphics.NewRect(collider.Width(), collider.Height())
			r.SetCentered(true)

			if en.HasComponent(components.Position) {
				position := components.Position.Get(en)
				r.Pos.SetBase(position.Vec)
			}

			r.SetFillColorScale(graphics.ColorScaleFromRGBA(0, 0, 0, 0))
			r.SetOutlineColorScale(graphics.RGB(0x0055ff))
			r.SetOutlineWidth(2)
			r.Draw(screen)
		}
	}
}
