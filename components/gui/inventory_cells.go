package gui

import (
	"github.com/ISMashtakov/mygame/items"
	"github.com/yohamta/donburi"
)

type SelectedItemData struct {
	CellNumber int
}

var SelectedCell = donburi.NewComponentType[SelectedItemData]()

type SelectCellRequestData struct {
	CellNumber int
}

var SelectCellRequest = donburi.NewComponentType[SelectCellRequestData]()

type LocationEnum int

const (
	InventoryLocation LocationEnum = iota
	DownPanelLocation
)

type CellLocation struct {
	CellNumber int
	Location   LocationEnum
}

var CellDragEvent = donburi.NewComponentType[CellLocation]()
var CellDropEvent = donburi.NewComponentType[CellLocation]()

type DragAndDropItemData struct {
	Item items.IItem
	From CellLocation
}

var DragAndDropItem = donburi.NewComponentType[DragAndDropItemData]()
