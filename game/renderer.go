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
	for en := range donburi.NewQuery(filter.Contains(components.Sprite)).Iter(world) {
		sprite := components.Sprite.Get(en)

		op := ebiten.DrawImageOptions{}

		if !sprite.Scale.IsZero() {
			op.GeoM.Scale(sprite.Scale.X, sprite.Scale.Y)
		}

		if en.HasComponent(components.Position) {
			position := components.Position.Get(en)
			op.GeoM.Translate(position.X, position.Y)
		}

		screen.DrawImage(sprite.Image, &op)
	}
}

func (r *Renderer) drawColliders(screen *ebiten.Image, world donburi.World) {
	if r.DrawColliders {
		for en := range donburi.NewQuery(filter.Contains(components.RectCollider)).Iter(world) {
			collider := components.RectCollider.Get(en)

			fmt.Println(collider.Rect)
			r := graphics.NewRect(collider.Width(), collider.Height())
			r.SetCentered(false)

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
