package actions

import "github.com/yohamta/donburi"

type ActionEnum int

const (
	HoeHit ActionEnum = iota + 1
	PickaxeHit
)

var Action = donburi.NewComponentType[ActionEnum]()

var ActionEnded = donburi.NewComponentType[ActionEnum]()
