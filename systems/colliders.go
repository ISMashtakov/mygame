package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type CollisionDetector struct {
}

func NewCollisionDetector() *CollisionDetector {
	return &CollisionDetector{}
}

func (d CollisionDetector) Update(world donburi.World) error {
	for d.fixCollision(world) > 0 {
	}
	return nil
}

func (d CollisionDetector) fixCollision(world donburi.World) int {
	countFixed := 0

	for en := range donburi.NewQuery(filter.Contains(components.Collider, components.MovementRequest)).Iter(world) {
		for en2 := range donburi.NewQuery(filter.Contains(components.Collider)).Iter(world) {
			if en == en2 {
				continue
			}

			rect := d.getRect(en)
			rect2 := d.getRect(en2)

			if rect.Intersects(rect2) {
				donburi.Remove[any](en, components.MovementRequest)
				countFixed++
				break
			}
		}

	}

	return countFixed
}

func (d CollisionDetector) getRect(en *donburi.Entry) gmath.Rect {
	collider := components.Collider.Get(en)

	offset := gmath.Vec{}
	if en.HasComponent(components.Position) {
		pos := components.Position.Get(en)
		offset = offset.Add(pos.Vec)
	}

	if en.HasComponent(components.MovementRequest) {
		moveReq := components.MovementRequest.Get(en)
		offset = offset.Add(moveReq.Vec)
	}

	return collider.Rect.Add(offset)
}
