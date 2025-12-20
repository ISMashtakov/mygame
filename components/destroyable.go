package components

import "github.com/yohamta/donburi"

type DestroyableData struct {
}

var Destroyable = donburi.NewComponentType[DestroyableData]()
