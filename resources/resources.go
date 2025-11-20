package resources

type ImageID = int

var (
	imageResources = map[ImageID]string{
		ImageCharacter: "resources/images/char.png",
		ImageGrass:     "resources/images/grass.png",
		ImageStone:     "resources/images/stone.png",
	}
)

const (
	ImageNone ImageID = iota
	ImageCharacter
	ImageGrass
	ImageStone
)
