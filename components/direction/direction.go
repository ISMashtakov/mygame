package direction

import "github.com/yohamta/donburi"

type DirectionEnum int

const (
	Up = iota + 1
	Right
	Down
	Left
)

var Direction = donburi.NewComponentType[DirectionEnum]()
