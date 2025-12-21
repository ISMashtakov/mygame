package systems

import (
	"time"

	"github.com/ISMashtakov/mygame/animations"
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/subsystems"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/ISMashtakov/mygame/utils/funcs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	PickaxeHitRequestHandlerCodename = "PickaxeHitRequestHandlerCodename"
)

type PickaxeHitRequestHandler struct {
	core.BaseSystem
	gardenCreator      background.GardenCreator
	collidersSubsystem subsystems.ColliderSearcher
	spriteCreator      entities.SimpeSpriteCreator
}

func NewPickaxeHitRequestHandler(spriteCreator entities.SimpeSpriteCreator) *PickaxeHitRequestHandler {
	return &PickaxeHitRequestHandler{
		BaseSystem: core.BaseSystem{
			Codename:        PickaxeHitRequestHandlerCodename,
			PreviousSystems: []string{AnimationCodename},
		},
		spriteCreator: spriteCreator,
	}
}

func (m *PickaxeHitRequestHandler) Update(world donburi.World) {
	for en := range don.IterByRequests(world, actions.PickaxeHitRequest) {
		hitArea := gmath.Vec{X: 10, Y: 10}

		rect := gmath.Rect{
			Min: en.Point.Sub(hitArea.Mulf(0.5)),
			Max: en.Point.Add(hitArea.Mulf(0.5)),
		}

		for _, destroyableEn := range m.collidersSubsystem.SearchByRect(world, rect, filter.Contains(components.Destroyable)) {
			m.destroyObj(world, destroyableEn)
		}
	}
}

func (m *PickaxeHitRequestHandler) destroyObj(world donburi.World, destroyableEn *donburi.Entry) {
	if destroyableEn.HasComponent(components.Position) && destroyableEn.HasComponent(components.Sprite) {
		pos := components.Position.Get(destroyableEn)
		sprite := components.Sprite.Get(destroyableEn)

		parts := m.spriteCreator.DivideSprite(world, *sprite, *pos)

		duration := time.Millisecond * 500
		getScaleAnimation := func(entry *donburi.Entry) *animations.ScaleAnimation {
			return animations.NewScaleAnimation(
				funcs.LineTo(duration.Seconds()*float64(ebiten.TPS()), -sprite.Image.Scale.X),
				funcs.LineTo(duration.Seconds()*float64(ebiten.TPS()), -sprite.Image.Scale.Y),
				components.Sprite.Get(entry),
			)
		}

		components.StartAnimation(
			world,
			*core.NewAnimationPlayer(
				duration,
				func() { world.Remove(parts[0].Entity()) },
				animations.NewMoveAnimation(funcs.Line(-0.3), funcs.SquareTo(8, -3, 50, 0), components.Position.Get(parts[0])),
				getScaleAnimation(parts[0]),
			),
		)
		components.StartAnimation(
			world,
			*core.NewAnimationPlayer(
				duration,
				func() { world.Remove(parts[1].Entity()) },
				animations.NewMoveAnimation(funcs.Line(0.3), funcs.SquareTo(8, -3, 50, 0), components.Position.Get(parts[1])),
				getScaleAnimation(parts[1]),
			),
		)
		components.StartAnimation(
			world,
			*core.NewAnimationPlayer(
				duration,
				func() { world.Remove(parts[2].Entity()) },
				animations.NewMoveAnimation(funcs.Line(0.3), funcs.SquareTo(8, 3, 50, 0), components.Position.Get(parts[2])),
				getScaleAnimation(parts[2]),
			),
		)
		components.StartAnimation(
			world,
			*core.NewAnimationPlayer(
				duration,
				func() { world.Remove(parts[3].Entity()) },
				animations.NewMoveAnimation(funcs.Line(-0.3), funcs.SquareTo(8, 3, 50, 0), components.Position.Get(parts[3])),
				getScaleAnimation(parts[3]),
			),
		)
	}

	world.Remove(destroyableEn.Entity())
}
