package actions

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/subsystems"
	"github.com/ISMashtakov/mygame/systems"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	SeedCodename = "SeedCodename"
)

type SeedHandler struct {
	core.BaseSystem

	collidersSubsystem subsystems.ColliderSearcher
	plantsCreator      *entities.PlantCreator
}

func NewSeedHandler(plantsCreator *entities.PlantCreator) *SeedHandler {
	return &SeedHandler{
		BaseSystem: core.BaseSystem{
			Codename:        SeedCodename,
			PreviousSystems: []string{systems.InputCodename},
			NextSystems:     []string{systems.PlantsCodename},
		},
		plantsCreator: plantsCreator,
	}
}

func (m *SeedHandler) Update(world donburi.World) {
	for seedData := range don.IterByRequests(world, actions.SeedRequest) {
		for _, garden := range m.collidersSubsystem.SearchByPoint(world, seedData.Point,
			filter.And(
				filter.Contains(components.Garden),
				filter.Not(filter.Contains(components.Filled)),
			)) {
			m.plantsCreator.Create(world, garden, seedData.Item)
			m.decSelectedItem(world)
		}
	}
}

func (m *SeedHandler) decSelectedItem(world donburi.World) {
	panelEntity, ok := donburi.NewQuery(filter.Contains(gui.SelectedCell, gui.DownPanel)).First(world)
	if !ok {
		panic("panel not found")
	}

	selectedCell := gui.SelectedCell.Get(panelEntity)
	downPanel := gui.DownPanel.Get(panelEntity)

	downPanel.AddItem(world, selectedCell.CellNumber, -1)
}
