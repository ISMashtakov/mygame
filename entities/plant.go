package entities

import (
	"time"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/constants/z"
	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/yohamta/donburi"
)

type PlantCreator struct {
	resourcesLoader resources.IResourceLoader
}

func NewPlantCreator(resourcesLoader resources.IResourceLoader) *PlantCreator {
	return &PlantCreator{
		resourcesLoader: resourcesLoader,
	}
}

func (c *PlantCreator) Create(world donburi.World, garden *donburi.Entry, item items.IItem) *donburi.Entry {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.RectCollider,
		components.Plant,
	)

	en := world.Entry(entity)

	components.Position.SetValue(en, components.PositionData{
		Vec: components.Position.Get(garden).Vec,
	})
	components.RectCollider.SetValue(en, *components.RectCollider.Get(garden))
	components.Sprite.SetValue(en, components.SpriteData{
		Z: z.PLANT,
	})
	components.Plant.SetValue(en, components.PlantData{
		SpriteSheet:  c.resourcesLoader.LoadSpriteSheet(resources.SheetCarrotPlant), // TODO: Придумать как по айтему это получать.
		GrowDuration: time.Second * 10,
		Stages:       4,
	})

	garden.AddComponent(components.Filled)

	return en
}
