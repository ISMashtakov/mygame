package game

import (
	"github.com/yohamta/donburi"
)

type ISystem interface {
	Update(world donburi.World)
}
