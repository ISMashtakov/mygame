package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/subsystems"
	"github.com/ISMashtakov/mygame/utils/filter2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	CollissionDetectorCodename = "collision_detector"
)

type CollisionDetector struct {
	core.BaseSystem
	colidersSubsystem subsystems.ColliderSearcher
}

func NewCollisionDetector() *CollisionDetector {
	return &CollisionDetector{
		BaseSystem: core.BaseSystem{
			Codename:        CollissionDetectorCodename,
			NextSystems:     []string{MovementCodename},
			PreviousSystems: []string{InputCodename},
		},
	}
}

func (d CollisionDetector) Update(world donburi.World) error {
	for d.fixCollision(world) > 0 {
	}
	return nil
}

func (d CollisionDetector) fixCollision(world donburi.World) int {
	countFixed := 0

	colliderFilter := filter2.ContainsAny(components.RectCollider, components.SpriteCollider)

	for en := range donburi.NewQuery(
		filter.And(
			filter.Contains(components.MovementRequest),
			colliderFilter,
		),
	).Iter(world) {
		for en2 := range donburi.NewQuery(
			filter.And(colliderFilter, filter.Contains(components.Obstacle)),
		).Iter(world) {
			if en == en2 {
				continue
			}

			if d.colidersSubsystem.IsIntersect(en, en2) {
				donburi.Remove[any](en, components.MovementRequest)
				countFixed++
				break
			}
		}
	}

	return countFixed
}
