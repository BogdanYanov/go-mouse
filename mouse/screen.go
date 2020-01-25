package mouse

// Screen height and width constants
const (
	Height uint32 = 768
	Width  uint32 = 1024
)

// Screen stores screen borders so that the mouse cursor does not run behind them
type Screen struct {
	width, height uint32
}

// NewScreen create new screen
func NewScreen(width, height uint32) *Screen {
	return &Screen{width: width, height: height}
}
