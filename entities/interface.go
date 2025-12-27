package entities

import (
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
)

type InterfaceCreator struct {
	itemsFactory *items.Factory
}

func NewInterfaceCreator(itemsFactory *items.Factory) *InterfaceCreator {
	return &InterfaceCreator{
		itemsFactory: itemsFactory,
	}
}

func (c InterfaceCreator) Create(world donburi.World) (*donburi.Entry, error) {
	entity := world.Create(
		gui.SelectedCell,
		gui.DownPanel,
		gui.Inventory,
	)

	entry := world.Entry(entity)

	// Стартовый выбор первой ячейки в нижней панели
	don.CreateRequest(world, gui.SelectCellRequest, &gui.SelectCellRequestData{})

	// TODO: Начальные предметы надо вынести в другое место

	panel := gui.DownPanel.Get(entry)
	panel.SetItem(world, 0, c.itemsFactory.Hoe())
	panel.SetItem(world, 1, c.itemsFactory.Pickaxe())

	return entry, nil
}
