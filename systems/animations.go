package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	AnimationCodename = "animation"
)

type Animation struct {
	core.BaseSystem

	characterCreator *entities.CharacterCreator
}

func NewAnimation(characterCreator *entities.CharacterCreator) *Animation {
	return &Animation{
		BaseSystem: core.BaseSystem{
			Codename:        AnimationCodename,
			PreviousSystems: []string{InputCodename},
		},
		characterCreator: characterCreator,
	}
}

func (s Animation) Update(world donburi.World) {
	for animationEntry := range donburi.NewQuery(filter.Contains(components.Animation)).Iter(world) {
		animation := components.Animation.Get(animationEntry)
		if animation.Player.Next() {
			components.DeleteAnimation(world, animationEntry)
		}
	}
}
