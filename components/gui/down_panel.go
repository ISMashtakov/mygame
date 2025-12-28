package gui

import (
	"fmt"

	"github.com/ISMashtakov/mygame/items"
	"github.com/ISMashtakov/mygame/utils/don"
	"github.com/yohamta/donburi"
)

type DownPanelData struct {
	items [9]items.IItem
}

func (d *DownPanelData) SetItem(world donburi.World, index int, item items.IItem) {
	d.validateIndex(index)

	d.items[index] = item

	don.Create(world, SetItemToGUIInventoryRequest, &SetItemToGUIInventoryRequestData{
		Location: CellLocation{
			CellNumber: index,
			Location:   DownPanelLocation,
		},
		Item: item,
	})
}

func (d *DownPanelData) AddItem(world donburi.World, index int, count int) {
	d.validateIndex(index)

	d.items[index].AddCount(count)

	don.Create(world, SetItemToGUIInventoryRequest, &SetItemToGUIInventoryRequestData{
		Location: CellLocation{
			CellNumber: index,
			Location:   DownPanelLocation,
		},
		Item: d.items[index],
	})
}

func (d *DownPanelData) GetItem(index int) items.IItem {
	d.validateIndex(index)

	return d.items[index]
}

func (d *DownPanelData) GetItems() []items.IItem {
	return d.items[:]
}

func (d *DownPanelData) validateIndex(index int) {
	if index >= len(d.items) || index < 0 {
		panic(fmt.Errorf("can't set items to index %d", index))
	}
}

var DownPanel = donburi.NewComponentType[DownPanelData]()
