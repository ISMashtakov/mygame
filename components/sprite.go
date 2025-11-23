package components

import (
	"github.com/ISMashtakov/mygame/core/images"
	"github.com/yohamta/donburi"
)

type SpriteData images.Image

var Sprite = donburi.NewComponentType[SpriteData]()
