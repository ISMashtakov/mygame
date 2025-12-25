package direction

import (
	"fmt"

	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type Enum int

const (
	Up = iota + 1
	Right
	Down
	Left
)

var Direction = donburi.NewComponentType[Enum]()

func GetDirectionVector(dir Enum) gmath.Vec {
	switch dir {
	case Up:
		return gmath.Vec{Y: -1}
	case Right:
		return gmath.Vec{X: 1}
	case Down:
		return gmath.Vec{Y: 1}
	case Left:
		return gmath.Vec{X: -1}
	}
	panic(fmt.Errorf("unknown direction: %d", dir))
}
