package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/utils/filter2"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/quasilyte/gmath"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	CollissionDetectorCodename = "collision_detector"
)

type collisionType struct {
	colliderType, colliderType2 donburi.IComponentType
}

type CollisionDetector struct {
	core.BaseSystem
}

func NewCollisionDetector() *CollisionDetector {
	return &CollisionDetector{
		core.BaseSystem{
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
		for en2 := range donburi.NewQuery(colliderFilter).Iter(world) {
			if en == en2 {
				continue
			}

			if d.isIntersect(en, en2) {
				donburi.Remove[any](en, components.MovementRequest)
				countFixed++
				break
			}
		}
	}

	return countFixed
}

func (d CollisionDetector) isIntersect(en, en2 *donburi.Entry) bool {
	collisionTypes := map[collisionType]func(en, en2 *donburi.Entry) bool{
		{components.RectCollider, components.RectCollider}:     d.isIntersectRectWithRect,
		{components.RectCollider, components.SpriteCollider}:   d.isIntersectRectWithSprite,
		{components.SpriteCollider, components.RectCollider}:   func(en, en2 *donburi.Entry) bool { return d.isIntersectRectWithSprite(en2, en) },
		{components.SpriteCollider, components.SpriteCollider}: d.isIntersectSpriteWithSprite,
	}

	for collisionType, isIntersectFunc := range collisionTypes {
		if en.HasComponent(collisionType.colliderType) && en2.HasComponent(collisionType.colliderType2) {
			if isIntersectFunc(en, en2) {
				return true
			}
		}
	}

	return false
}

func (d CollisionDetector) isIntersectRectWithRect(en, en2 *donburi.Entry) bool {
	rect := d.getRect(en)
	rect2 := d.getRect(en2)

	return rect.Intersects(rect2)
}

func (d CollisionDetector) getRect(en *donburi.Entry) gmath.Rect {
	collider := components.RectCollider.Get(en)

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

func (d CollisionDetector) isIntersectSpriteWithSprite(en, en2 *donburi.Entry) bool {
	sprite1, sprite2 := components.Sprite.Get(en), components.Sprite.Get(en2)

	var pos1, pos2 gmath.Vec
	scale1 := lo.Ternary(sprite1.Scale.IsZero(), gmath.Vec{X: 1, Y: 1}, sprite1.Scale)
	scale2 := lo.Ternary(sprite2.Scale.IsZero(), gmath.Vec{X: 1, Y: 1}, sprite2.Scale)

	if en.HasComponent(components.Position) {
		pos := components.Position.Get(en)
		pos1 = pos.Vec
	}
	if en2.HasComponent(components.Position) {
		pos := components.Position.Get(en2)
		pos2 = pos.Vec
	}

	bounds1 := gmath.RectFromStd(sprite1.Image.Bounds())
	bounds2 := gmath.RectFromStd(sprite2.Image.Bounds())

	imageSize1 := gmath.Vec{X: bounds1.Width(), Y: bounds1.Height()}
	imageSize2 := gmath.Vec{X: bounds2.Width(), Y: bounds2.Height()}

	rect1 := gmath.Rect{
		Min: pos1,
		Max: pos1.Add(imageSize1.Mul(scale1)),
	}
	rect2 := gmath.Rect{
		Min: pos2,
		Max: pos2.Add(imageSize2.Mul(scale2)),
	}

	imageRec1 := rect1.ToStd()
	imageRec2 := rect2.ToStd()

	// Сначала проверяем пересечение bounding box
	if !imageRec1.Overlaps(imageRec2) {
		return false
	}

	// Находим область пересечения
	overlap := imageRec1.Intersect(imageRec2)

	// Проверяем каждый пиксель в области пересечения
	for y := overlap.Min.Y; y < overlap.Max.Y; y++ {
		for x := overlap.Min.X; x < overlap.Max.X; x++ {
			relVec := gmath.Vec{X: float64(x), Y: float64(y)}
			rel1 := relVec.Sub(pos1).Div(scale1)
			rel2 := relVec.Sub(pos2).Div(scale2)
			// Проверяем, что оба пикселя непрозрачны
			_, _, _, a1 := render.AtImage(sprite1.Image, rel1).RGBA()
			_, _, _, a2 := render.AtImage(sprite2.Image, rel2).RGBA()
			if a1 > 0 && a2 > 0 {
				return true
			}
		}
	}

	return false
}

func (d CollisionDetector) isIntersectRectWithSprite(en, en2 *donburi.Entry) bool {
	rect1 := d.getRect(en)
	sprite2 := components.Sprite.Get(en2)

	var pos2 gmath.Vec
	scale2 := lo.Ternary(sprite2.Scale.IsZero(), gmath.Vec{X: 1, Y: 1}, sprite2.Scale)

	if en2.HasComponent(components.Position) {
		pos := components.Position.Get(en2)
		pos2 = pos.Vec
	}

	bounds2 := gmath.RectFromStd(sprite2.Image.Bounds())

	imageSize2 := gmath.Vec{X: bounds2.Width(), Y: bounds2.Height()}

	rect2 := gmath.Rect{
		Min: pos2,
		Max: pos2.Add(imageSize2.Mul(scale2)),
	}

	imageRec1 := rect1.ToStd()
	imageRec2 := rect2.ToStd()

	// Сначала проверяем пересечение bounding box
	if !imageRec1.Overlaps(imageRec2) {
		return false
	}

	// Находим область пересечения
	overlap := imageRec1.Intersect(imageRec2)

	// Проверяем каждый пиксель в области пересечения
	for y := overlap.Min.Y; y < overlap.Max.Y; y++ {
		for x := overlap.Min.X; x < overlap.Max.X; x++ {
			rel2 := gmath.Vec{X: float64(x), Y: float64(y)}.Sub(pos2).Div(scale2)
			_, _, _, a2 := sprite2.Image.At(int(rel2.X), int(rel2.Y)).RGBA()

			if a2 > 0 {
				return true
			}
		}
	}

	return false
}
