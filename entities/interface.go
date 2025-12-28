package entities

import (
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/gui/guicomponents"
	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
)

type InterfaceCreator struct {
	itemsFactory *items.Factory
	inventory    *guicomponents.Inventory
	downPanel    *guicomponents.DownPanel
}

func NewInterfaceCreator(
	itemsFactory *items.Factory,
	inventory *guicomponents.Inventory,
	downPanel *guicomponents.DownPanel,
) *InterfaceCreator {
	return &InterfaceCreator{
		itemsFactory: itemsFactory,
		inventory:    inventory,
		downPanel:    downPanel,
	}
}

func (c InterfaceCreator) Create(world donburi.World) *donburi.Entry {
	entity := world.Create(
		gui.SelectedCell,
		gui.DownPanel,
		gui.Inventory,
	)

	entry := world.Entry(entity)

	// Стартовый выбор первой ячейки в нижней панели
	don.Create(world, gui.SelectCellRequest, &gui.SelectCellRequestData{})

	// TODO: Начальные предметы надо вынести в другое место

	panel := gui.DownPanel.Get(entry)
	panel.SetItem(world, 0, c.itemsFactory.Hoe())
	panel.SetItem(world, 1, c.itemsFactory.Pickaxe())

	carrotSeeds := c.itemsFactory.CarrotSeed()
	carrotSeeds.AddCount(10)
	panel.SetItem(world, 2, carrotSeeds)

	c.setOnClickToItemCells(world)

	return entry
}

func (c InterfaceCreator) setOnClickToItemCells(world donburi.World) {
	registerHandlers := func(cell *guicomponents.InventoryCell, index int, location gui.LocationEnum) {
		cell.SetOnDrag(func() {
			don.Create(world, gui.CellDragEvent, &gui.CellLocation{
				CellNumber: index,
				Location:   location,
			})
		})
		cell.SetOnDrop(func() {
			don.Create(world, gui.CellDropEvent, &gui.CellLocation{
				CellNumber: index,
				Location:   location,
			})
		})
	}

	for i, cell := range c.inventory.Cells() {
		registerHandlers(cell, i, gui.InventoryLocation)
	}

	for i, cell := range c.downPanel.Cells() {
		registerHandlers(cell, i, gui.DownPanelLocation)
	}
}
