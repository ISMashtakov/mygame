package subsystems

import (
	"iter"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/utils/filter2"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type collisionType struct {
	colliderType, colliderType2 donburi.IComponentType
}

type ColliderSearcher struct{}

func NewColliderSearcher() *ColliderSearcher {
	return &ColliderSearcher{}
}

func (s ColliderSearcher) SearchByEntry(
	world donburi.World,
	entry *donburi.Entry,
	filters ...filter.LayoutFilter,
) []*donburi.Entry {
	var result []*donburi.Entry

	for en := range s.iterByColliders(world, filters...) {
		if s.IsIntersect(en, entry) {
			result = append(result, en)
		}
	}

	return result
}

func (s ColliderSearcher) SearchByRect(
	world donburi.World,
	rect gmath.Rect,
	filters ...filter.LayoutFilter,
) []*donburi.Entry {
	var result []*donburi.Entry

	for en := range s.iterByColliders(world, filters...) {
		if en.HasComponent(components.RectCollider) && s.getRect(en).Intersects(rect) {
			result = append(result, en)
		}

		if en.HasComponent(components.SpriteCollider) && s.isIntersectRectWithSprite(rect, en) {
			result = append(result, en)
		}
	}

	return result
}

func (s ColliderSearcher) SearchByPoint(
	world donburi.World,
	point gmath.Vec,
	filters ...filter.LayoutFilter,
) []*donburi.Entry {
	var result []*donburi.Entry

	for en := range s.iterByColliders(world, filters...) {
		if en.HasComponent(components.RectCollider) && s.isInRect(en, point) {
			result = append(result, en)
		}
		if en.HasComponent(components.SpriteCollider) && s.isInSprite(en, point) {
			result = append(result, en)
		}
	}

	return result
}

func (s ColliderSearcher) iterByColliders(
	world donburi.World,
	filters ...filter.LayoutFilter,
) iter.Seq[*donburi.Entry] {
	colliderFilter := filter2.ContainsAny(components.RectCollider, components.SpriteCollider)
	notDisabledFilter := filter.Not(filter.Contains(components.DisabledColliders))
	return donburi.NewQuery(filter.And(append(filters, colliderFilter, notDisabledFilter)...)).Iter(world)
}

func (s ColliderSearcher) IsIntersect(en, en2 *donburi.Entry) bool {
	if en.HasComponent(components.DisabledColliders) || en2.HasComponent(components.DisabledColliders) {
		return false
	}

	collisionTypes := map[collisionType]func(en, en2 *donburi.Entry) bool{
		{components.RectCollider, components.RectCollider}:     s.isIntersectRectWithRect,
		{components.RectCollider, components.SpriteCollider}:   func(en, en2 *donburi.Entry) bool { return s.isIntersectRectWithSprite(s.getRect(en), en2) },
		{components.SpriteCollider, components.RectCollider}:   func(en, en2 *donburi.Entry) bool { return s.isIntersectRectWithSprite(s.getRect(en2), en) },
		{components.SpriteCollider, components.SpriteCollider}: s.isIntersectSpriteWithSprite,
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

func (s ColliderSearcher) getRect(en *donburi.Entry) gmath.Rect {
	collider := components.RectCollider.Get(en)

	offset := gmath.Vec{}
	if en.HasComponent(components.Position) {
		pos := components.Position.Get(en)
		offset = offset.Add(pos.Vec)
	}

	if en.HasComponent(components.Movement) {
		moveReq := components.Movement.Get(en)
		offset = offset.Add(moveReq.Vec)
	}

	return collider.Rect.Add(offset)
}

func (s ColliderSearcher) isInRect(en *donburi.Entry, point gmath.Vec) bool {
	rect := s.getRect(en)

	return rect.Contains(point)
}

func (s ColliderSearcher) isInSprite(en *donburi.Entry, point gmath.Vec) bool {
	sprite := components.Sprite.Get(en)

	var pos gmath.Vec

	if en.HasComponent(components.Position) {
		pos = components.Position.Get(en).Vec
	}

	bounds := gmath.RectFromStd(sprite.Image.Bounds())

	imageSize := gmath.Vec{X: bounds.Width(), Y: bounds.Height()}.Mul(sprite.Image.Scale)

	rel2 := point.
		Add(imageSize.Mulf(0.5)).
		Sub(pos).
		Div(sprite.Image.Scale)
	_, _, _, a2 := render.AtImage(sprite.Image.Image, rel2).RGBA()

	return a2 > 0
}

func (s ColliderSearcher) isIntersectRectWithRect(en, en2 *donburi.Entry) bool {
	rect := s.getRect(en)
	rect2 := s.getRect(en2)

	return rect.Intersects(rect2)
}

func (s ColliderSearcher) isIntersectSpriteWithSprite(en, en2 *donburi.Entry) bool {
	sprite1, sprite2 := components.Sprite.Get(en), components.Sprite.Get(en2)

	var pos1, pos2 gmath.Vec

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

	imageSize1 := gmath.Vec{X: bounds1.Width(), Y: bounds1.Height()}.Mul(sprite1.Image.Scale)
	imageSize2 := gmath.Vec{X: bounds2.Width(), Y: bounds2.Height()}.Mul(sprite2.Image.Scale)

	imageRec1 := s.getSpriteRect(en, imageSize1, pos1).ToStd()
	imageRec2 := s.getSpriteRect(en2, imageSize2, pos2).ToStd()

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
			rel1 := relVec.Add(imageSize1.Mulf(0.5)).
				Sub(pos1).
				Div(sprite1.Image.Scale)
			rel2 := relVec.Add(imageSize2.Mulf(0.5)).
				Sub(pos2).
				Div(sprite2.Image.Scale)
			// Проверяем, что оба пикселя непрозрачны
			_, _, _, a1 := render.AtImage(sprite1.Image.Image, rel1).RGBA()
			_, _, _, a2 := render.AtImage(sprite2.Image.Image, rel2).RGBA()
			if a1 > 0 && a2 > 0 {
				return true
			}
		}
	}

	return false
}

func (s ColliderSearcher) isIntersectRectWithSprite(rect gmath.Rect, en2 *donburi.Entry) bool {
	sprite2 := components.Sprite.Get(en2)

	var pos2 gmath.Vec

	if en2.HasComponent(components.Position) {
		pos := components.Position.Get(en2)
		pos2 = pos.Vec
	}

	bounds2 := gmath.RectFromStd(sprite2.Image.Bounds())

	imageSize2 := gmath.Vec{X: bounds2.Width(), Y: bounds2.Height()}
	imageSize2 = imageSize2.Mul(sprite2.Image.Scale)

	imageRec1 := rect.ToStd()
	imageRec2 := s.getSpriteRect(en2, imageSize2, pos2).ToStd()

	// Сначала проверяем пересечение bounding box
	if !imageRec1.Overlaps(imageRec2) {
		return false
	}

	// Находим область пересечения
	overlap := imageRec1.Intersect(imageRec2)

	// Проверяем каждый пиксель в области пересечения
	for y := overlap.Min.Y; y < overlap.Max.Y; y++ {
		for x := overlap.Min.X; x < overlap.Max.X; x++ {
			rel2 := gmath.Vec{X: float64(x), Y: float64(y)}.
				Add(imageSize2.Mulf(0.5)).
				Sub(pos2).
				Div(sprite2.Image.Scale)
			_, _, _, a2 := render.AtImage(sprite2.Image.Image, rel2).RGBA()

			if a2 > 0 {
				return true
			}
		}
	}

	return false
}

func (s ColliderSearcher) getSpriteRect(spriteEntity *donburi.Entry, imageSize, pos gmath.Vec) gmath.Rect {
	spriteCollider := components.SpriteCollider.Get(spriteEntity)

	if spriteCollider.ActiveZone != nil {
		return gmath.Rect{
			Min: spriteCollider.ActiveZone.Min.Add(pos),
			Max: spriteCollider.ActiveZone.Max.Add(pos),
		}
	}

	return gmath.Rect{
		Min: pos.Sub(imageSize.Mulf(0.5)),
		Max: pos.Add(imageSize.Mulf(0.5)),
	}
}
