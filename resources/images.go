package resources

type ImageID = int

var (
	imageResources = map[ImageID]string{
		ImageCharacterMoving:     "resources/images/character/char.png",
		ImageCharacterHoeHitting: "resources/images/character/hoe_hitting.png",
		ImageGrass:               "resources/images/grass.png",
		ImageStone:               "resources/images/stone.png",
	}
)

const (
	ImageNone ImageID = iota

	ImageCharacterMoving
	ImageCharacterHoeHitting

	ImageGrass
	ImageStone
)
