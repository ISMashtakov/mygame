package components

import (
	"github.com/ISMashtakov/mygame/items"
	"github.com/yohamta/donburi"
)

type DestroyableData struct {
	Resources []items.IItem
}

var Destroyable = donburi.NewComponentType[DestroyableData]()
