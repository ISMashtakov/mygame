package components

import (
	"github.com/ISMashtakov/mygame/items"
	"github.com/yohamta/donburi"
)

type PropData struct {
	Item items.IItem
}

var Prop = donburi.NewComponentType[PropData]()
