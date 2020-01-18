package mouse

type Screen struct {
	width, height uint32
}

func NewScreen(width, height uint32) *Screen {
	return &Screen{width:width, height:height}
}
