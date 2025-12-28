package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/gui/guicomponents"
	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/subsystems"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	PropsTakingCodename = "PropsTakingCodename"
)

type PropsTaking struct {
	core.BaseSystem
	subsystems.ColliderSearcher

	inventory *guicomponents.Inventory
	downPanel *guicomponents.DownPanel
}

func NewPropsTaking(inventory *guicomponents.Inventory, downPanel *guicomponents.DownPanel) *PropsTaking {
	return &PropsTaking{
		BaseSystem: core.BaseSystem{
			Codename: PropsTakingCodename,
		},
		inventory: inventory,
		downPanel: downPanel,
	}
}

func (s PropsTaking) Update(world donburi.World) {
	charEntry, ok := donburi.NewQuery(filter.Contains(components.Character)).First(world)
	if !ok {
		return
	}

	for _, propEntry := range s.ColliderSearcher.SearchByEntry(world, charEntry, filter.Contains(components.Prop, components.RectCollider)) {
		prop := components.Prop.Get(propEntry)
		if s.addItemToInventory(world, prop.Item) {
			propEntry.Remove()
		}
	}
}

func (s PropsTaking) addItemToInventory(world donburi.World, item items.IItem) bool {
	panel := don.GetComponent(world, gui.DownPanel)
	inventory := don.GetComponent(world, gui.Inventory)

	for index, cell := range panel.GetItems() {
		if cell != nil && cell.GetType() == item.GetType() && cell.GetCount() < cell.GetMaxStackSize() {
			panel.AddItem(world, index, 1)
			return true
		}
	}

	for index, cell := range inventory.GetItems() {
		if cell != nil && cell.GetType() == item.GetType() && cell.GetCount() < cell.GetMaxStackSize() {
			inventory.AddItem(world, index, 1)
			return true
		}
	}

	for index, cell := range panel.GetItems() {
		if cell == nil {
			panel.SetItem(world, index, item)
			return true
		}
	}

	for index, cell := range inventory.GetItems() {
		if cell == nil {
			inventory.SetItem(world, index, item)
			return true
		}
	}

	return false
}
