package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	MovementCodename = "movement"
)

type Movement struct {
	core.BaseSystem
}

func NewMovement() *Movement {
	return &Movement{
		core.BaseSystem{
			Codename:        MovementCodename,
			PreviousSystems: []string{InputCodename, CollissionDetectorCodename},
		},
	}
}

func (m *Movement) Update(world donburi.World) error {
	for en := range donburi.NewQuery(filter.Contains(components.Position, components.MovementRequest)).Iter(world) {
		pos, moveRequest := components.Position.Get(en), components.MovementRequest.Get(en)

		pos.Vec = pos.Vec.Add(moveRequest.Vec)

		donburi.Remove[any](en, components.MovementRequest)
	}

	return nil
}
