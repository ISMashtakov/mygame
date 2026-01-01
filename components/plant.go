package components

import (
	"time"

	"github.com/ISMashtakov/mygame/core/images"
	"github.com/quasilyte/gmath"
	"github.com/yohamta/donburi"
)

type PlantData struct {
	SpriteSheet       *images.SpritesSheet
	Stages            int
	GrowDuration      time.Duration
	CurrentFrame      int
	Offset            gmath.Vec
	IsReadyForHarvest bool
}

var Plant = donburi.NewComponentType[PlantData]()

var ReadyForHarvest = donburi.NewTag()
