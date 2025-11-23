package components

import (
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image images.Image
	Z     float32
}

func (d SpriteData) Order() int {
	return int(d.Z)
}

var Sprite = donburi.NewComponentType[SpriteData]()
