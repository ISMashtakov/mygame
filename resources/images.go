package resources

type ImageID = int

var (
	imageResources = map[ImageID]string{
		ImageCharacterMoving:     "resources/images/character/walk.png",
		ImageCharacterHoeHitting: "resources/images/character/hit.png",
		ImageGrass:               "resources/images/grass.png",
		ImageStone:               "resources/images/stone.png",
		ImageGarden:              "resources/images/garden.png",
	}
)

const (
	ImageNone ImageID = iota

	ImageCharacterMoving
	ImageCharacterHoeHitting

	ImageGrass
	ImageStone
	ImageGarden
)
