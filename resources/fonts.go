package resources

type FontID = int

var (
	fontResources = map[FontID]string{
		FontSimple: "resources/fonts/notosans-regular.ttf",
	}
)

const (
	FontSimple ImageID = iota
)
