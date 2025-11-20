package systems

import (
	"fmt"

	"github.com/ISMashtakov/mygame/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Movement struct {
}

func NewMovement() *Movement {
	return &Movement{}
}

func (m *Movement) Update(world donburi.World) error {
	for en := range donburi.NewQuery(filter.Contains(components.Position, components.MovementRequest)).Iter(world) {
		pos, moveRequest := components.Position.Get(en), components.MovementRequest.Get(en)

		pos.Vec = pos.Vec.Add(moveRequest.Vec)
		components.Position.Set(en, pos)
		fmt.Println(pos)

		components.MovementRequest.Set(en, nil)
	}

	return nil
}
