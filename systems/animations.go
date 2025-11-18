package systems

import (
	"image"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core/direction"
	"github.com/ISMashtakov/mygame/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

const (
	sizeX float64 = 17
	sizeY float64 = 25
)

type SwapSpriteByWalkingAnimation struct {
	TPS                 int
	framesInAnimation   int
	animationLoopLength int
	resourceLoader      resources.IResourceLoader
}

func NewSwapSpriteByWalkingAnimation(TPS int, resourceLoader resources.IResourceLoader) *SwapSpriteByWalkingAnimation {
	return &SwapSpriteByWalkingAnimation{
		TPS:                 TPS,
		framesInAnimation:   4,
		animationLoopLength: int(float64(TPS) * 0.5),
		resourceLoader:      resourceLoader,
	}
}

func (s SwapSpriteByWalkingAnimation) Update(world donburi.World) error {
	for en := range donburi.NewQuery(filter.Contains(components.Sprite, components.Character, components.WalkingAnimation, components.Speed)).Iter(world) {
		walkingAnimation, speed := components.WalkingAnimation.Get(en), components.Speed.Get(en)

		if !speed.IsZero() {
			switch {
			case speed.Y > 0:
				walkingAnimation.Direction = direction.Down
			case speed.Y < 0:
				walkingAnimation.Direction = direction.Up
			case speed.X > 0:
				walkingAnimation.Direction = direction.Right
			case speed.X < 0:
				walkingAnimation.Direction = direction.Left
			}
		}

		walkingAnimation.Frame += 1

		if walkingAnimation.Frame >= s.animationLoopLength || speed.IsZero() {
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
		components.Sprite.SetValue(en, components.SpriteData{
			Image: subImage.(*ebiten.Image),
			Scale: &gmath.Vec{X: sizeX / float64(rect.Dx()), Y: sizeY / float64(rect.Dy())},
		})

	}

	return nil
}
