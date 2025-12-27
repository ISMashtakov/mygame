package gui

import (
	"fmt"

	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
)

type InventoryData struct {
	items [4 * 4]items.IItem
}

func (d *InventoryData) SetItem(world donburi.World, index int, item items.IItem) {
	d.validateIndex(index)

	d.items[index] = item

	don.CreateRequest(world, SetItemToGUIInventoryRequest, &SetItemToGUIInventoryRequestData{
		Index: index,
		Item:  item,
	})
}

func (d *InventoryData) GetItem(index int) items.IItem {
	d.validateIndex(index)

	return d.items[index]
}

func (d *InventoryData) GetItems() []items.IItem {
	return d.items[:]
}

func (d *InventoryData) validateIndex(index int) {
	if index >= len(d.items) || index < 0 {
		panic(fmt.Errorf("can't set items to index %d", index))
	}
}

var Inventory = donburi.NewComponentType[InventoryData]()

type SetItemToGUIInventoryRequestData struct {
	Index int
	Item  items.IItem
}

var SetItemToGUIInventoryRequest = donburi.NewComponentType[SetItemToGUIInventoryRequestData]()

var SwitchInventaryStatusRequest = donburi.NewTag()
