package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/constants"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/subsystems"
	"github.com/ISMashtakov/mygame/utils"
	"github.com/ISMashtakov/mygame/utils/filter2"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	HoeHitCheckerCodename = "hoe_hit_checker"
)

type HoeHitChecker struct {
	core.BaseSystem
	gardenCreator      background.GardenCreator
	collidersSubsystem subsystems.ColliderSearcher
}

func NewHoeHitChecker(garderCreator background.GardenCreator) *HoeHitChecker {
	return &HoeHitChecker{
		BaseSystem: core.BaseSystem{
			Codename:        HoeHitCheckerCodename,
			PreviousSystems: []string{AnimationCodename},
		},
		gardenCreator: garderCreator,
	}
}

func (m *HoeHitChecker) Update(world donburi.World) error {
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

			donburi.Remove[any](en, actions.ActionEnded)
		}

	}

	return nil
}
