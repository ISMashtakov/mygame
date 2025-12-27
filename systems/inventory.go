package systems

import (
	"fmt"

	"github.com/ISMashtakov/mygame/components/gui"
	"github.com/ISMashtakov/mygame/core"
	guicomponents "github.com/ISMashtakov/mygame/gui/guicomponents"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
)

const (
	InventoryCodename = "InventoryCodename"
)

type Inventory struct {
	core.BaseSystem

	inventory *guicomponents.Inventory
}

func NewInventory(inventory *guicomponents.Inventory) *Inventory {
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

	inventorySize := len(c.inventory.Cells())
	for request := range don.IterByRequests(world, gui.SetItemToGUIInventoryRequest) {
		if request.Index >= inventorySize*inventorySize || request.Index < 0 {
			panic(fmt.Errorf("invalid value of set item cell %d", request.Index))
		}

		c.inventory.Cells()[request.Index/inventorySize][request.Index%inventorySize].SetImage(request.Item.GetImage())
	}
}
