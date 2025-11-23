package entities

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/constants"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type CharacterCreator struct {
	TargetImageSize gmath.Vec
}

func NewCharacterCreator() *CharacterCreator {
	return &CharacterCreator{
		TargetImageSize: gmath.Vec{X: 17, Y: 25},
	}
}

func (c CharacterCreator) Create(world donburi.World) (donburi.Entity, error) {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.MovementRequest,
		components.Character,
		components.WalkingAnimation,
		// Подумать над спрайтовым коллайдером, но проблема, что анимация меняет спрайт и можно застрять в текстуре.
		components.RectCollider,
		direction.Direction,
	)

	en := world.Entry(entity)

	rect := gmath.Rect{
		Min: constants.CharacterColliderSize.Mulf(-0.5).Add(gmath.Vec{Y: 2}),
		Max: constants.CharacterColliderSize.Mulf(0.5),
	}

	components.RectCollider.SetValue(en, components.RectColliderData{Rect: rect})
	direction.Direction.SetValue(en, direction.Down)

	return entity, nil
}
