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

func (m *Movement) Update(world donburi.World) {
	for en := range donburi.NewQuery(filter.Contains(components.Position, components.Movement)).Iter(world) { //TODO: тут тоже
		pos, moveRequest := components.Position.Get(en), components.Movement.Get(en)

		pos.Vec = pos.Vec.Add(moveRequest.Vec)

		donburi.Remove[any](en, components.Movement)
	}
}
