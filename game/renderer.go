package game

import (
	"fmt"
	"slices"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

var ()

type Renderer struct {
	DrawColliders bool
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Draw(screen *ebiten.Image, world donburi.World) {
	r.drawSprites(screen, world)
	r.drawColliders(screen, world)
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("FPS: %0.2f, TPS: %0.2f, entities: %d", ebiten.ActualFPS(), ebiten.ActualTPS(), world.Len()),
	)
}

func (r *Renderer) drawSprites(screen *ebiten.Image, world donburi.World) {
	spritesEntity := slices.Collect(
		donburi.NewQuery(filter.Contains(components.Sprite)).Iter(world),
	)

	slices.SortFunc(spritesEntity, r.drawOrderFunc)

	cameraEn, ok := donburi.NewQuery(filter.Contains(components.Position, components.Camera)).First(world)
	if !ok {
		panic("can not find camera")
	}

	cameraPos := components.Position.Get(cameraEn)

	for _, en := range spritesEntity {
		sprite := components.Sprite.Get(en)

		op := ebiten.DrawImageOptions{}

		op.GeoM.Scale(sprite.Image.Scale.X, sprite.Image.Scale.Y)

		op.GeoM.Translate(
			-float64(sprite.Image.Bounds().Dx())/2*sprite.Image.Scale.X,
			-float64(sprite.Image.Bounds().Dy())/2*sprite.Image.Scale.Y,
		)

		if sprite.Image.Flip {
			op.GeoM.Scale(-1, 1)
		}

		if en.HasComponent(components.Position) {
			position := components.Position.Get(en)
			op.GeoM.Translate(position.X, position.Y)
		}
		r.applyCameraMods(&op.GeoM, *cameraPos)

		screen.DrawImage(sprite.Image.Image, &op)
	}
}

func (r *Renderer) applyCameraMods(geom *ebiten.GeoM, cameraPos components.PositionData) {
	geom.Translate(-cameraPos.Vec.X, -cameraPos.Vec.Y)

	width, height := ebiten.WindowSize()
	windowSize := gmath.Vec{X: float64(width), Y: float64(height)}

	shift := windowSize.Mulf(0.5).Mul(constants.TargetLayout.Div(windowSize))
	geom.Translate(shift.X, shift.Y)
}

func (r *Renderer) drawOrderFunc(en1, en2 *donburi.Entry) int {
	sprite1 := components.Sprite.Get(en1)
	sprite2 := components.Sprite.Get(en2)

	if sprite1.Z < sprite2.Z {
		return -1
	} else if sprite1.Z > sprite2.Z {
		return 1
	}

	var position1 components.PositionData
	var position2 components.PositionData

	if en1.HasComponent(components.Position) {
		position1 = *components.Position.Get(en1)
	}
	if en2.HasComponent(components.Position) {
		position2 = *components.Position.Get(en2)
	}

	return int(position1.Y - position2.Y)
}

func (r *Renderer) drawColliders(screen *ebiten.Image, world donburi.World) {
	if r.DrawColliders {
		for en := range donburi.NewQuery(filter.Contains(components.RectCollider)).Iter(world) {
			collider := components.RectCollider.Get(en)

			r := graphics.NewRect(collider.Width(), collider.Height())
			r.SetCentered(false)
			r.Pos.Set(nil, collider.Min.X, collider.Min.Y)

			if en.HasComponent(components.Position) {
				position := components.Position.Get(en)
				r.Pos.SetBase(position.Vec)
			}

			r.SetFillColorScale(graphics.ColorScaleFromRGBA(0, 0, 0, 0))
			r.SetOutlineColorScale(graphics.RGB(0x0055ff))
			r.SetOutlineWidth(2)
			r.Draw(screen)
		}

		for en := range donburi.NewQuery(filter.Contains(components.SpriteCollider)).Iter(world) {
			collider := components.SpriteCollider.Get(en)
			if collider.ActiveZone == nil {
				continue
			}

			r := graphics.NewRect(collider.ActiveZone.Width(), collider.ActiveZone.Height())
			r.SetCentered(false)
			r.Pos.Set(nil, collider.ActiveZone.Min.X, collider.ActiveZone.Min.Y)

			if en.HasComponent(components.Position) {
				position := components.Position.Get(en)
				r.Pos.SetBase(position.Vec)
			}

			r.SetFillColorScale(graphics.ColorScaleFromRGBA(0, 0, 0, 0))
			r.SetOutlineColorScale(graphics.RGB(0x5522ff))
			r.SetOutlineWidth(2)
			r.Draw(screen)
		}
	}
}
