package resources

type ImageID = int

var (
	imageResources = map[ImageID]string{
		ImageCharacterMoving:     "resources/images/character/walk.png",
		ImageCharacterHoeHitting: "resources/images/character/hit.png",
		ImageGrass:               "resources/images/grass.png",
		ImageCoal:                "resources/images/coal.png",
		ImageGarden:              "resources/images/garden.png",
		// Items
		ImageItemHoe:     "resources/images/items/hoe.png",
		ImageItemPickaxe: "resources/images/items/pickaxe.png",
		ImageItemCoal:    "resources/images/items/coal.png",
		// Seeds
		ImageItemSeedCarrot: "resources/images/items/seeds/carrot.png",
	}
)

const (
	ImageNone ImageID = iota

	ImageCharacterMoving
	ImageCharacterHoeHitting

	ImageGrass
	ImageCoal
	ImageGarden

	// ITEMS.

	ImageItemHoe
	ImageItemPickaxe
	ImageItemCoal

	// SEEDS
	ImageItemSeedCarrot
)
