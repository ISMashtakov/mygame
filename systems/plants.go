package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/entities/background"
	"github.com/ISMashtakov/mygame/utils"
	"github.com/ISMashtakov/mygame/utils/funcs"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	PlantsCodename = "PlantsCodename"
)

type Plants struct {
	core.BaseSystem
	gardenCreator *background.GardenCreator
}

func NewPlants(gardenCreator *background.GardenCreator) *Plants {
	return &Plants{
		BaseSystem: core.BaseSystem{
			Codename:        PlantsCodename,
			PreviousSystems: []string{InputCodename},
		},
		gardenCreator: gardenCreator,
	}
}

func (p *Plants) Update(world donburi.World) {
	for entity := range donburi.NewQuery(
		filter.And(
			filter.Contains(components.Plant, components.Sprite),
			filter.Not(filter.Contains(components.ReadyForHarvest)),
		),
	).Iter(world) {
		plant, sprite := components.Plant.Get(entity), components.Sprite.Get(entity)
		gropDuration := funcs.X(plant.GrowDuration)

		if plant.CurrentFrame >= int(gropDuration-1) {
			entity.AddComponent(components.ReadyForHarvest)
			sprite.Image = plant.SpriteSheet.Get(plant.Stages-1, 0)
			sprite.Image.Scale = render.GetImageScale(sprite.Image.Bounds(), p.gardenCreator.TargetImageSize)
			continue
		}
		plant.CurrentFrame += 1

		sprite.Image = plant.SpriteSheet.Get(utils.GetStep(float64(plant.CurrentFrame), gropDuration, plant.Stages-1), 0)
		sprite.Image.Scale = render.GetImageScale(sprite.Image.Bounds(), p.gardenCreator.TargetImageSize)
	}
}
