package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/constants"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/subsystems"
	"github.com/ISMashtakov/mygame/utils"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/ISMashtakov/mygame/utils/filter2"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

const (
	GardenCreatingRequestHandlerCodename = "GardenCreatingRequestHandlerCodename"
)

type GardenCreatingRequestHandler struct {
	core.BaseSystem

	gardenCreator      background.GardenCreator
	collidersSubsystem subsystems.ColliderSearcher
}

func NewGardenCreatingRequestHandler(garderCreator background.GardenCreator) *GardenCreatingRequestHandler {
	return &GardenCreatingRequestHandler{
		BaseSystem: core.BaseSystem{
			Codename:        GardenCreatingRequestHandlerCodename,
			PreviousSystems: []string{AnimationCodename},
		},
		gardenCreator: garderCreator,
	}
}

func (m *GardenCreatingRequestHandler) Update(world donburi.World) {
	for en := range don.IterByRequests(world, actions.GardenCreatingRequest) {
		point := utils.FloorByNearestStepVec(en.Point, constants.TileSize)

		rect := gmath.Rect{
			Min: point.Sub(m.gardenCreator.TargetImageSize.Mulf(0.5)),
			Max: point.Add(m.gardenCreator.TargetImageSize.Mulf(0.5)),
		}

		if len(
			m.collidersSubsystem.SearchByRect(world, rect, filter2.ContainsAny(components.Garden, components.Obstacle)),
		) == 0 {
			m.gardenCreator.Create(world, components.PositionData{Vec: point})
		}
	}
}
