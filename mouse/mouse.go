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
	leftButton  button
	rightButton button
	sensitivity uint8
	wheel       wheel
}

// NewMouse create a new mouse.
func NewMouse(screen Screen) *Mouse {
	return &Mouse{
		screen.width / 2,
		screen.height / 2,
		button{},
		button{},
		minSettingValue,
		wheel{defaultSettingValue},
	}
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

// Sensitivity sets a new value for mouse sensitivity.
func (m *Mouse) Sensitivity(val uint8) {
	if val > maxSettingValue {
		m.sensitivity = maxSettingValue
	} else if val == 0 {
		m.sensitivity = minSettingValue
	} else {
		m.sensitivity = val
	}
}

// LeftButtonDown simulates the left button pressed.
func (m *Mouse) LeftButtonDown() {
	m.leftButton.down()
}

// RightButtonDown simulates the right button pressed.
func (m *Mouse) RightButtonDown() {
	m.rightButton.down()
}

// LeftButtonUp simulates the left button released.
func (m *Mouse) LeftButtonUp() {
	m.leftButton.up()
}

// RightButtonUp simulates the right button released.
func (m *Mouse) RightButtonUp() {
	m.rightButton.up()
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
		m.posX, m.posY, m.sensitivity, m.leftButton.isPressed, m.rightButton.isPressed, m.wheel.scrollValue)
}

// Reset returns default settings and mouse states.
func (m *Mouse) Reset(screen Screen) {
	m.posX = screen.width / 2
	m.posY = screen.height / 2
	m.LeftButtonUp()
	m.RightButtonUp()
	m.wheel.scrollValue = defaultSettingValue
	m.Sensitivity(minSettingValue)
}
