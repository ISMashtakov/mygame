package resources

import (
	"github.com/quasilyte/gmath"
)

type SpriteSheetID = int

type spriteSheetData struct {
	path     string
	cellSize gmath.Vec
}

var (
	spriteSheetResources = map[SpriteSheetID]spriteSheetData{
		SheetCarrotPlant: {
			path:     "resources/images/items/seeds/carrot_plant_atlas.png",
			cellSize: gmath.Vec{X: 64, Y: 80},
		},
	}
)

const (
	SheetCarrotPlant SpriteSheetID = iota
)

//64 80
