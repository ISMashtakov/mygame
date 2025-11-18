package core

import (
	"github.com/yohamta/donburi"
)

type ISystem interface {
	Update(world donburi.World) error
}
