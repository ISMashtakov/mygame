package systems

import (
	"fmt"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/components/direction"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/ISMashtakov/mygame/entities"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	AnimationCodename = "animation"
)

type SwapSpriteByWalkingAnimation struct {
	core.BaseSystem
	characterCreator *entities.CharacterCreator
	animation        *images.Animation
}

func NewSwapSpriteByWalkingAnimation(resourceLoader resources.IResourceLoader, characterCreator *entities.CharacterCreator) (*SwapSpriteByWalkingAnimation, error) {
	animation, err := resourceLoader.LoadAnimation(resources.AnimationCharacterWalking)
	if err != nil {
		return nil, fmt.Errorf("can't get animation: %w", err)
	}

	return &SwapSpriteByWalkingAnimation{
		BaseSystem: core.BaseSystem{
			Codename:        AnimationCodename,
			PreviousSystems: []string{InputCodename},
		},
		characterCreator: characterCreator,
		animation:        animation,
	}, nil
}

func (s SwapSpriteByWalkingAnimation) Update(world donburi.World) error {
	for en := range donburi.NewQuery(
		filter.Contains(
			components.Sprite, components.Character, components.WalkingAnimation, direction.Direction,
		)).Iter(world) {
		dir := direction.Direction.GetValue(en)

		if !en.HasComponent(components.MovementRequest) {
			s.animation.Reset()
		}

		subImage := s.animation.Next(dir)
		sprite := components.SpriteData{
			Image: subImage,
			Scale: render.GetImageScale(subImage.Bounds(), s.characterCreator.TargetImageSize),
		}

		components.Sprite.SetValue(en, sprite)
	}

	return nil
}
