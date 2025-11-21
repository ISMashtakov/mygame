package entities

import (
	"image"

	"github.com/ISMashtakov/mygame/components"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type CharacterCreator struct {
	TargetImageSize image.Rectangle
}

func NewCharacterCreator() *CharacterCreator {
	return &CharacterCreator{
		TargetImageSize: image.Rect(0, 0, 17, 25),
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
	)

	en := world.Entry(entity)

	components.RectCollider.SetValue(en, components.RectColliderData{Rect: gmath.RectFromStd(c.TargetImageSize)})

	return entity, nil
}
