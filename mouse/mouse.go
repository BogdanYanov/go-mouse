package mouse

import (
	"fmt"
)

const (
	minSettingValue     uint8 = 1
	maxSettingValue     uint8 = 10
	defaultSettingValue uint8 = 5
)

// Mouse stores mouse states and settings.
type Mouse struct {
	posX        uint32
	posY        uint32
	leftBtn     Button
	rightBtn    Button
	sensitivity uint8
	wheel       Wheel
}

// NewMouse create a new mouse.
func NewMouse(screen Screen) *Mouse {
	mouse := &Mouse{
		screen.width / 2,
		screen.height / 2,
		Button{},
		Button{},
		minSettingValue,
		Wheel{defaultSettingValue},
	}
	return mouse
}

// Move moving the mouse cursor to x, y coordinates.
func (m *Mouse) Move(x, y uint32, screen Screen) {
	if x > screen.width {
		x = screen.width
	}
	if y > screen.height {
		y = screen.height
	}
Loop:
	for {
		if x > m.posX {
			if m.posX+uint32(m.sensitivity) <= x {
				m.posX += uint32(m.sensitivity)
			} else {
				m.posX += x - m.posX
			}
		} else {
			if m.posX-uint32(m.sensitivity) >= x {
				m.posX -= uint32(m.sensitivity)
			} else {
				m.posX -= m.posX - x
			}
		}

		if y > m.posY {
			if m.posY+uint32(m.sensitivity) <= y {
				m.posY += uint32(m.sensitivity)
			} else {
				m.posY += y - m.posY
			}
		} else {
			if m.posY-uint32(m.sensitivity) >= y {
				m.posY -= uint32(m.sensitivity)
			} else {
				m.posY -= m.posY - y
			}
		}

		if m.posX == x && m.posY == y {
			break Loop
		}
	}
}

// sensitivity sets a new value for mouse sensitivity when moving.
func (m *Mouse) SetSensitivity(val uint8) {
	if val > maxSettingValue {
		m.sensitivity = maxSettingValue
	} else if val == 0 {
		m.sensitivity = minSettingValue
	} else {
		m.sensitivity = val
	}
}

// leftBtnDown simulates the left button pressed.
func (m *Mouse) LeftBtnDown() {
	m.leftBtn.down()
}

// rightBtnDown simulates the right button pressed.
func (m *Mouse) RightBtnDown() {
	m.rightBtn.down()
}

// leftBtnUp simulates the left button released.
func (m *Mouse) LeftBtnUp() {
	m.leftBtn.up()
}

// rightBtnUp simulates the right button released.
func (m *Mouse) RightBtnUp() {
	m.rightBtn.up()
}

// ScrollUp simulates mouse scroll up.
func (m *Mouse) ScrollUp() {
	m.wheel.scrollUp()
}

// ScrollDown simulates mouse scroll down.
func (m *Mouse) ScrollDown() {
	m.wheel.scrollDown()
}

// Info displays mouse states and settings.
func (m *Mouse) Info() {
	fmt.Printf("Mouse information:\n"+
		"X position - %d\n"+
		"Y position - %d\n"+
		"Sensitivity - %d\n"+
		"Is left button pressed? - %v\n"+
		"Is right button pressed? - %v\n"+
		"Scroll value - %d\n",
		m.posX, m.posY, m.sensitivity, m.leftBtn.btnPressed, m.rightBtn.btnPressed, m.wheel.scrollValue)
}

// Reset returns default settings and mouse states.
func (m *Mouse) Reset(screen Screen) {
	m.posX = screen.width / 2
	m.posY = screen.height / 2
	m.LeftBtnUp()
	m.RightBtnUp()
	m.wheel.scrollValue = defaultSettingValue
	m.SetSensitivity(minSettingValue)
}
