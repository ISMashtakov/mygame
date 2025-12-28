package entities

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/constants/z"
	"github.com/ISMashtakov/mygame/utils"
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

func (c CharacterCreator) Create(world donburi.World) *donburi.Entry {
	entity := world.Create(
		components.Position,
		components.Sprite,
		components.Movement,
		components.Character,
		components.WalkingAnimation,
		// Подумать над спрайтовым коллайдером, но проблема, что анимация меняет спрайт и можно застрять в текстуре.
		components.RectCollider,
		direction.Direction,
		components.CurrentAnimation,
	)

	en := world.Entry(entity)

	rect := utils.GetRectOfBottomOfParent(c.TargetImageSize, 0.5)

	components.RectCollider.SetValue(en, components.RectColliderData{Rect: rect})
	direction.Direction.SetValue(en, direction.Down)
	components.Sprite.Get(en).Z = z.OBJ

	return en
}
