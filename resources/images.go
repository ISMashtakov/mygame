package resources

type ImageID = int

var (
	imageResources = map[ImageID]string{
		ImageCharacterMoving:     "resources/images/character/walk.png",
		ImageCharacterHoeHitting: "resources/images/character/hit.png",
		ImageGrass:               "resources/images/grass.png",
		ImageCoil:                "resources/images/coil.png",
		ImageGarden:              "resources/images/garden.png",
		// HOE
		ImageItemHoe:     "resources/images/items/hoe.png",
		ImageItemPickaxe: "resources/images/items/pickaxe.png",
	}
)

const (
	ImageNone ImageID = iota

	ImageCharacterMoving
	ImageCharacterHoeHitting

	ImageGrass
	ImageCoil
	ImageGarden

	// ITEMS
	ImageItemHoe
	ImageItemPickaxe
)
