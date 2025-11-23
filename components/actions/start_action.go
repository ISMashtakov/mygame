package actions

import "github.com/yohamta/donburi"

type ActionEnum int

var (
	HoeHit ActionEnum = 1
)

var Action = donburi.NewComponentType[ActionEnum]()
