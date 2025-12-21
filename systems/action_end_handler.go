package systems

import (
	"time"

	"github.com/ISMashtakov/mygame/animations"
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/constants"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/subsystems"
	"github.com/ISMashtakov/mygame/utils"
	"github.com/ISMashtakov/mygame/utils/filter2"
	"github.com/ISMashtakov/mygame/utils/funcs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	ActionEndHandlerCodename = "action_end_handler"
)

type ActionEndHandler struct {
	core.BaseSystem
	gardenCreator      background.GardenCreator
	collidersSubsystem subsystems.ColliderSearcher
	spriteCreator      entities.SimpeSpriteCreator
}

func NewHoeHitChecker(garderCreator background.GardenCreator, spriteCreator entities.SimpeSpriteCreator) *ActionEndHandler {
	return &ActionEndHandler{
		BaseSystem: core.BaseSystem{
			Codename:        ActionEndHandlerCodename,
			PreviousSystems: []string{AnimationCodename},
		},
		gardenCreator: garderCreator,
		spriteCreator: spriteCreator,
	}
}

func (m *ActionEndHandler) Update(world donburi.World) {
	for en := range donburi.NewQuery(filter.Contains(actions.ActionEnded, direction.Direction, components.Position)).Iter(world) {
		action, dir, position := actions.ActionEnded.Get(en), direction.Direction.Get(en), components.Position.Get(en)
		switch *action {
		case actions.HoeHit:
			point := position.Vec.Add(direction.GetDirectionVector(*dir).Mul(constants.TileSize))

			// сдвиг для красоты
			if *dir != direction.Down {
				point.Y += 10
			}

			point = utils.FloorByNearestStepVec(point, constants.TileSize)

			rect := gmath.Rect{
				Min: point.Sub(m.gardenCreator.TargetImageSize.Mulf(0.5)),
				Max: point.Add(m.gardenCreator.TargetImageSize.Mulf(0.5)),
			}

			if len(m.collidersSubsystem.SearchByRect(world, rect, filter2.ContainsAny(components.Garden, components.Obstacle))) == 0 {
				m.gardenCreator.Create(world, components.PositionData{Vec: point})
			}
		case actions.PickaxeHit:
			point := position.Vec.Add(direction.GetDirectionVector(*dir).Mul(constants.TileSize))

			// сдвиг для красоты
			if *dir != direction.Down {
				point.Y += 10
			}

			hitArea := gmath.Vec{X: 10, Y: 10}

			rect := gmath.Rect{
				Min: point.Sub(hitArea.Mulf(0.5)),
				Max: point.Add(hitArea.Mulf(0.5)),
			}

			for _, destroyableEn := range m.collidersSubsystem.SearchByRect(world, rect, filter.Contains(components.Destroyable)) {
				m.destroyObj(world, destroyableEn)
			}
		}

		donburi.Remove[any](en, actions.ActionEnded)

	}
}

func (m *ActionEndHandler) destroyObj(world donburi.World, destroyableEn *donburi.Entry) {
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
