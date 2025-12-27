package systems

import (
	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/core"
	guicomponents "github.com/ISMashtakov/mygame/gui/components"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
)

const (
	InventoryCodename = "InventoryCodename"
)

type Inventory struct {
	core.BaseSystem

	inventory guicomponents.Inventory
}

func NewInventory(inventory guicomponents.Inventory) *Inventory {
	return &Inventory{
		BaseSystem: core.BaseSystem{
			Codename:        InventoryCodename,
			PreviousSystems: []string{InputCodename},
		},
		inventory: inventory,
	}
}

func (c *Inventory) Update(world donburi.World) {
	for range don.IterByRequests(world, gui.SwitchInventaryStatusRequest) {
		c.inventory.Switch()
	}
}
