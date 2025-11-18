package physics

import (
	"github.com/ISMashtakov/mygame/core"
	"github.com/mlange-42/ark/ecs"
	"github.com/quasilyte/gmath"
)

type Movement struct {
	speedFilter *ecs.Filter2[core.Position, Speed]
}

func NewMovement(world *ecs.World) *Movement {
	return &Movement{
		speedFilter: ecs.NewFilter2[core.Position, Speed](world),
	}
}

func (m *Movement) Update() error {
	query := m.speedFilter.Query()
	for query.Next() {
		pos, speed := query.Get()

		pos.Vec = pos.Vec.Add(speed.Vec)
		speed.Vec = gmath.Vec{}
	}

	return nil
}
