package gui

import "github.com/yohamta/donburi"

type SelectedItemData struct {
	CellNumber int
}

var SelectedCell = donburi.NewComponentType[SelectedItemData]()

type SelectCellRequestData struct {
	CellNumber int
}

var SelectCellRequest = donburi.NewComponentType[SelectCellRequestData]()
