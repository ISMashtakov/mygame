package systems

import (
	"fmt"

	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/core"
	guicomponents "github.com/ISMashtakov/mygame/gui/guicomponents"
	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
)

const (
	InterfaceCodename = "InterfaceCodename"
)

type Inventory struct {
	core.BaseSystem

	inventory *guicomponents.Inventory
	downPanel *guicomponents.DownPanel
}

func NewInventory(
	inventory *guicomponents.Inventory,
	downPanel *guicomponents.DownPanel,
) *Inventory {
	return &Inventory{
		BaseSystem: core.BaseSystem{
			Codename:        InterfaceCodename,
			PreviousSystems: []string{InputCodename},
		},
		inventory: inventory,
		downPanel: downPanel,
	}
}

func (c *Inventory) Update(world donburi.World) {
	c.handleSwitchInventoryStatus(world)

	c.handleSetItemToInventoryRequest(world)

	c.handleCellClickedRequests(world)

	c.updateSelecting(world)
}

func (c *Inventory) handleSetItemToInventoryRequest(world donburi.World) {
	for request := range don.IterByRequests(world, gui.SetItemToGUIInventoryRequest) {
		cells := c.getCellsByLocation(request.Location.Location)

		size := len(cells)
		if request.Location.CellNumber >= size*size || request.Location.CellNumber < 0 {
			panic(fmt.Errorf("invalid value of set item cell %d", request.Location.CellNumber))
		}

		cell := cells[request.Location.CellNumber]
		if request.Item == nil {
			cell.SetImage(nil)
			cell.SetCount(0)
			continue
		}

		cell.SetImage(request.Item.GetImage())
		if request.Item.GetMaxStackSize() > 1 {
			cell.SetCount(request.Item.GetCount())
		}
	}
}

func (c *Inventory) getCellsByLocation(location gui.LocationEnum) []*guicomponents.InventoryCell {
	switch location {
	case gui.DownPanelLocation:
		return c.downPanel.Cells()
	case gui.InventoryLocation:
		return c.inventory.Cells()
	}

	panic(fmt.Errorf("unknown location %v", location))
}

func (c *Inventory) handleSwitchInventoryStatus(world donburi.World) {
	for range don.IterByRequests(world, gui.SwitchInventaryStatusRequest) {
		c.inventory.Switch()
	}
}

func (c *Inventory) updateSelecting(world donburi.World) {
	cell := don.GetComponent(world, gui.SelectedCell)

	for request := range don.IterByRequests(world, gui.SelectCellRequest) {
		if cell.CellNumber >= len(c.downPanel.Cells()) || cell.CellNumber < 0 {
			panic(fmt.Errorf("invalid value of selected cell %d", cell.CellNumber))
		}

		if request.CellNumber >= len(c.downPanel.Cells()) || request.CellNumber < 0 {
			panic(fmt.Errorf("invalid value of requested selected cell %d", request.CellNumber))
		}

		c.downPanel.Cells()[cell.CellNumber].Disable()
		c.downPanel.Cells()[request.CellNumber].Enable()

		cell.CellNumber = request.CellNumber
	}
}

func (c *Inventory) handleCellClickedRequests(world donburi.World) {
	for location := range don.IterByRequests(world, gui.CellDragEvent) {
		item := c.getItemByLocation(world, *location)
		if item == nil {
			continue
		}

		don.DeleteAllEntitiesWithComponent(world, gui.DragAndDropItem)
		don.Create(world, gui.DragAndDropItem, &gui.DragAndDropItemData{
			Item: item,
			From: *location,
		})
	}

	for location := range don.IterByRequests(world, gui.CellDropEvent) {
		dragAndDrop := don.GetComponent(world, gui.DragAndDropItem)

		fromItem := c.getItemByLocation(world, dragAndDrop.From)
		toItem := c.getItemByLocation(world, *location)

		if fromItem != nil && toItem != nil && fromItem.GetType() == toItem.GetType() {
			freeSlots := toItem.GetMaxStackSize() - toItem.GetCount()
			forDrag := min(freeSlots, fromItem.GetCount())
			c.addItemByLocation(world, dragAndDrop.From, -forDrag)
			c.addItemByLocation(world, *location, forDrag)
		} else {
			c.setItemByLocation(world, *location, fromItem)
			c.setItemByLocation(world, dragAndDrop.From, toItem)
		}
	}
}

func (c *Inventory) getItemByLocation(world donburi.World, location gui.CellLocation) items.IItem {
	switch location.Location {
	case gui.DownPanelLocation:
		downPanel := don.GetComponent(world, gui.DownPanel)
		return downPanel.GetItem(location.CellNumber)
	case gui.InventoryLocation:
		inventory := don.GetComponent(world, gui.Inventory)
		return inventory.GetItem(location.CellNumber)
	}

	return nil
}

func (c *Inventory) setItemByLocation(world donburi.World, location gui.CellLocation, item items.IItem) {
	switch location.Location {
	case gui.DownPanelLocation:
		downPanel := don.GetComponent(world, gui.DownPanel)
		downPanel.SetItem(world, location.CellNumber, item)
	case gui.InventoryLocation:
		inventory := don.GetComponent(world, gui.Inventory)
		inventory.SetItem(world, location.CellNumber, item)
	}
}

func (c *Inventory) addItemByLocation(world donburi.World, location gui.CellLocation, count int) {
	switch location.Location {
	case gui.DownPanelLocation:
		downPanel := don.GetComponent(world, gui.DownPanel)
		downPanel.AddItem(world, location.CellNumber, count)
	case gui.InventoryLocation:
		inventory := don.GetComponent(world, gui.Inventory)
		inventory.AddItem(world, location.CellNumber, count)
	}
}
