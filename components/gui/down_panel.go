package gui

import (
	"github.com/ISMashtakov/mygame/items"
	"github.com/yohamta/donburi"
)

type DownPanelData struct {
	Items [9]items.IItem
}

var DownPanel = donburi.NewComponentType[DownPanelData]()

type SetItemToDownPanelRequestData struct {
	Index int
	Item  items.IItem
}

var SetItemToDownPanelRequest = donburi.NewComponentType[SetItemToDownPanelRequestData]()
