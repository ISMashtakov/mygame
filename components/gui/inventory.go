package gui

import (
	"github.com/ISMashtakov/mygame/items"
	"github.com/yohamta/donburi"
)

type InventoryData struct {
	cellStorage
}

func NewInventoryData() *InventoryData {
	return &InventoryData{
		cellStorage{
			size:     4 * 4,
			items:    make([]items.IItem, 4*4),
			location: InventoryLocation,
		},
	}
}

var Inventory = donburi.NewComponentType[InventoryData]()

type SetItemToGUIInventoryRequestData struct {
	Location CellLocation
	Item     items.IItem
}

var SetItemToGUIInventoryRequest = donburi.NewComponentType[SetItemToGUIInventoryRequestData]()

var SwitchInventaryStatusRequest = donburi.NewTag()
