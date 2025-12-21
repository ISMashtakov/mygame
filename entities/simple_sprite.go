package entities

import (
	"image"

	"github.com/ISMashtakov/mygame/components"
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type SimpeSpriteCreator struct{}

func NewSimpeSpriteCreator() *SimpeSpriteCreator {
	return &SimpeSpriteCreator{}
}

func (c SimpeSpriteCreator) Create(world donburi.World, sprite components.SpriteData, pos components.PositionData) donburi.Entity {
	entity := world.Create(
		components.Position,
		components.Sprite,
	)

	entry := world.Entry(entity)

	components.Sprite.Set(entry, &sprite)
	components.Position.Set(entry, &pos)

	return entity
}

// DivideSprite Разделяет справйт на 4 и возвращает в порядке по часовой стрелке начиная с верхнего левого
func (c SimpeSpriteCreator) DivideSprite(world donburi.World, sprite components.SpriteData, pos components.PositionData) []*donburi.Entry {
	getSprite := func(imageRect image.Rectangle, shiftX, shiftY float64) *donburi.Entry {
		ent := c.Create(world, components.SpriteData{Z: sprite.Z, Image: images.Image{
			Scale: sprite.Image.Scale,
			Flip:  sprite.Image.Flip,
			Image: sprite.Image.Image.SubImage(imageRect).(*ebiten.Image),
		}}, components.PositionData{
			Vec: pos.Add(
				gmath.VecFromStd(sprite.Image.Bounds().Size()).Mul(sprite.Image.Scale).Mulf(0.25).Mul(gmath.Vec{X: shiftX, Y: shiftY}),
			),
		})

		return world.Entry(ent)
	}

	if sprite.Image.Flip {
		lu := getSprite(image.Rect(sprite.Image.Bounds().Dx()/2, 0, sprite.Image.Bounds().Dx(), sprite.Image.Bounds().Dy()/2), -1, -1)
		ru := getSprite(image.Rect(0, 0, sprite.Image.Bounds().Dx()/2, sprite.Image.Bounds().Dy()/2), 1, -1)
		rd := getSprite(image.Rect(0, sprite.Image.Bounds().Dy()/2, sprite.Image.Bounds().Dx()/2, sprite.Image.Bounds().Dy()), 1, 1)
		ld := getSprite(image.Rect(sprite.Image.Bounds().Dx()/2, sprite.Image.Bounds().Dy()/2, sprite.Image.Bounds().Dx(), sprite.Image.Bounds().Dy()), -1, 1)

		return []*donburi.Entry{lu, ru, rd, ld}
	} else {
		lu := getSprite(image.Rect(0, 0, sprite.Image.Bounds().Dx()/2, sprite.Image.Bounds().Dy()/2), -1, -1)
		ru := getSprite(image.Rect(sprite.Image.Bounds().Dx()/2, 0, sprite.Image.Bounds().Dx(), sprite.Image.Bounds().Dy()/2), 1, -1)
		rd := getSprite(image.Rect(sprite.Image.Bounds().Dx()/2, sprite.Image.Bounds().Dy()/2, sprite.Image.Bounds().Dx(), sprite.Image.Bounds().Dy()), 1, 1)
		ld := getSprite(image.Rect(0, sprite.Image.Bounds().Dy()/2, sprite.Image.Bounds().Dx()/2, sprite.Image.Bounds().Dy()), -1, 1)

		return []*donburi.Entry{lu, ru, rd, ld}
	}
}
