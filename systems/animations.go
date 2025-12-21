package systems

import (
	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/actions"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/constants/z"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	AnimationCodename = "animation"
)

type SwapSpriteByAnimation struct {
	core.BaseSystem
	characterCreator    *entities.CharacterCreator
	walkingAnimation    *images.Animation
	hoeHittingAnimation *images.Animation
}

func NewSwapSpriteByAnimation(resourceLoader resources.IResourceLoader, characterCreator *entities.CharacterCreator) *SwapSpriteByAnimation {
	walkingAnimation := resourceLoader.LoadAnimation(resources.AnimationCharacterWalking)
	hoeHittingAnimation := resourceLoader.LoadAnimation(resources.AnimationCharacterHoeHitting)

	return &SwapSpriteByAnimation{
		BaseSystem: core.BaseSystem{
			Codename:        AnimationCodename,
			PreviousSystems: []string{InputCodename},
		},
		characterCreator: characterCreator,

		walkingAnimation:    walkingAnimation,
		hoeHittingAnimation: hoeHittingAnimation,
	}
}

func (s SwapSpriteByAnimation) Update(world donburi.World) {
	for en := range donburi.NewQuery(
		filter.And(
			filter.Contains(components.Sprite, components.Character, components.WalkingAnimation, direction.Direction),
		)).Iter(world) {
		dir := direction.Direction.GetValue(en)

		if !en.HasComponent(components.Movement) {
			s.walkingAnimation.Reset()
		}

		var subImage images.Image

		if en.HasComponent(actions.Action) {
			action := actions.Action.Get(en)
			subImage = s.updateAction(en, *action, dir)
		} else {
			subImage = s.walkingAnimation.Next(dir)
		}

		components.Sprite.SetValue(en, components.SpriteData{
			Image: subImage,
			Z:     z.OBJ,
		})
	}

	s.newAnimationUpdate(world)
}

func (s SwapSpriteByAnimation) newAnimationUpdate(world donburi.World) {
	for animationEntry := range donburi.NewQuery(filter.Contains(components.Animation)).Iter(world) {
		animation := components.Animation.Get(animationEntry)
		if animation.Player.Next() {
			components.DeleteAnimation(world, animationEntry)
		}
	}
}

func (s SwapSpriteByAnimation) updateAction(character *donburi.Entry, action actions.ActionEnum, dir direction.DirectionEnum) images.Image {
	var subImage images.Image
	switch action {
	case actions.HoeHit:
		subImage = s.hoeHittingAnimation.Next(dir)

		if s.hoeHittingAnimation.IsFinish() {
			s.hoeHittingAnimation.Reset()
			donburi.Remove[any](character, actions.Action)
			donburi.Add(character, actions.ActionEnded, &action)
		}
	case actions.PickaxeHit:
		subImage = s.hoeHittingAnimation.Next(dir)

		if s.hoeHittingAnimation.IsFinish() {
			s.hoeHittingAnimation.Reset()
			donburi.Remove[any](character, actions.Action)
			donburi.Add(character, actions.ActionEnded, &action)
		}
	}

	return subImage
}
