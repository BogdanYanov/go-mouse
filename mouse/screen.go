package mouse

const (
	width, height uint32 = 1024, 768
)

type Screen struct {
	width, height uint32
}

func NewScreen() *Screen {
	return &Screen{width: width, height: height}
}
