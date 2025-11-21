package systems

import (
	"image"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core"
	"github.com/ISMashtakov/mygame/core/direction"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/ISMashtakov/mygame/utils/render"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	AnimationCodename = "animation"
)

type SwapSpriteByWalkingAnimation struct {
	core.BaseSystem
	TPS                 int
	framesInAnimation   int
	animationLoopLength int
	resourceLoader      resources.IResourceLoader

	TargetImageSize image.Rectangle
}

func NewSwapSpriteByWalkingAnimation(TPS int, resourceLoader resources.IResourceLoader) *SwapSpriteByWalkingAnimation {
	return &SwapSpriteByWalkingAnimation{
		BaseSystem: core.BaseSystem{
			Codename:        AnimationCodename,
			PreviousSystems: []string{InputCodename},
		},
		TPS:                 TPS,
		framesInAnimation:   4,
		animationLoopLength: int(float64(TPS) * 0.5),
		resourceLoader:      resourceLoader,
		TargetImageSize:     image.Rect(0, 0, 17, 25),
	}
}

func (s SwapSpriteByWalkingAnimation) Update(world donburi.World) error {
	for en := range donburi.NewQuery(filter.Contains(components.Sprite, components.Character, components.WalkingAnimation, components.MovementRequest)).Iter(world) {
		walkingAnimation, moveReq := components.WalkingAnimation.Get(en), components.MovementRequest.Get(en)

		if !moveReq.IsZero() {
			switch {
			case moveReq.Y > 0:
				walkingAnimation.Direction = direction.Down
			case moveReq.Y < 0:
				walkingAnimation.Direction = direction.Up
			case moveReq.X > 0:
				walkingAnimation.Direction = direction.Right
			case moveReq.X < 0:
				walkingAnimation.Direction = direction.Left
			}
		}

		walkingAnimation.Frame += 1

		if walkingAnimation.Frame >= s.animationLoopLength || moveReq.IsZero() {
			walkingAnimation.Frame = 0
		}

		left, up, right, down := 80, 60, 250, 310

		switch walkingAnimation.Direction {
		case direction.Up:
			up += 320 * 3
			down += 320 * 3
		case direction.Right:
			up += 320 * 2
			down += 320 * 2
		case direction.Down:
		case direction.Left:
			up += 320 * 1
			down += 320 * 1
		}

		frame := int((float64(walkingAnimation.Frame) / float64(s.animationLoopLength)) * 4)

		left += 320 * frame
		right += 320 * frame

		rect := image.Rect(left, up, right, down)

		resImage, err := s.resourceLoader.LoadImage(resources.ImageCharacter)
		if err != nil {
			return err
		}

		subImage := resImage.SubImage(rect)
		sprite := components.SpriteData{
			Image: subImage.(*ebiten.Image),
			Scale: render.GetImageScale(rect, s.TargetImageSize),
		}

		components.Sprite.SetValue(en, sprite)
		components.Collider.SetValue(en, *components.GetColliderDataBySprite(sprite))
	}

	return nil
}
