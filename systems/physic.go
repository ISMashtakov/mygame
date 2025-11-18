package systems

import (
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
	for en := range donburi.NewQuery(filter.Contains(components.Position, components.Speed)).Iter(world) {
		pos, speed := components.Position.Get(en), components.Speed.Get(en)

		pos.Vec = pos.Vec.Add(speed.Vec)

		components.Speed.Set(en, &components.SpeedData{})
		components.Position.Set(en, pos)
	}

	return nil
}
