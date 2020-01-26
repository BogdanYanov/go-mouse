package mouse

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

var testInfoStr = `Mouse information:
X position - 512
Y position - 384
Sensitivity - 5
Is left button pressed? - true
Is right button pressed? - false
Scroll value - 9
`

func TestMouse_New(t *testing.T) {
	s := NewScreen(Width, Height)
	m := NewMouse(*s, &Button{}, &Button{}, NewWheel())
	if m == nil {
		t.Errorf("mouse is nil")
	}
}

func TestMouse_Move(t *testing.T) {
	var x, y uint32 = 1260, 1024
	var width, height uint32 = 1024, 768
	s := NewScreen(width, height)
	m := NewMouse(*s, &Button{}, &Button{}, NewWheel())
	m.Move(x, y, *s)
	if m.posX != s.width || m.posY != s.height {
		t.Errorf("Coords out of screen. Expected X - %d, Y - %d; Get X - %d, Y - %d",
			s.width,
			s.height,
			m.posX,
			m.posY)
	}
	m.Sensitivity(10)
	m.Move(968, 743, *s)
	if m.posX != 968 || m.posY != 743 {
		t.Errorf("Coords out of screen. Expected X - %d, Y - %d; Get X - %d, Y - %d",
			968,
			743,
			m.posX,
			m.posY)
	}
	m.Move(1260, 1024, *s)
	if m.posX != s.width || m.posY != s.height {
		t.Errorf("Coords out of screen. Expected X - %d, Y - %d; Get X - %d, Y - %d",
			s.width,
			s.height,
			m.posX,
			m.posY)
	}
}

func TestMouse_Sensitivity(t *testing.T) {
	s := NewScreen(Width, Height)
	m := NewMouse(*s, &Button{}, &Button{}, NewWheel())
	m.Sensitivity(11)
	if m.sensitivity != maxSettingValue {
		t.Errorf("Sensitivity out of range. Expected sensitivity - %d, get - %d", maxSettingValue, m.sensitivity)
	}
	m.Sensitivity(0)
	if m.sensitivity != minSettingValue {
		t.Errorf("Sensitivity out of range. Expected sensitivity - %d, get - %d", minSettingValue, m.sensitivity)
	}
}

func TestMouse_WheelScrollUp(t *testing.T) {
	s := NewScreen(Width, Height)
	m := NewMouse(*s, &Button{}, &Button{}, NewWheel())
	scrollValue := m.wheel.State()
	m.ScrollUp()
	if m.wheel.State() != scrollValue + 1 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", scrollValue + 1, m.wheel.State())
	}
	for i := m.wheel.State(); i < maxSettingValue + 2; i++ {
		m.ScrollUp()
	}
	if m.wheel.State() != maxSettingValue {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", maxSettingValue, m.wheel.State())
	}
}

func TestMouse_WheelScrollDown(t *testing.T) {
	s := NewScreen(Width, Height)
	m := NewMouse(*s, &Button{}, &Button{}, NewWheel())
	scrollValue := m.wheel.State()
	m.ScrollDown()
	if m.wheel.State() != scrollValue - 1 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", scrollValue - 1, m.wheel.State())
	}
	for i := m.wheel.State(); i > minSettingValue - 1; i-- {
		m.ScrollDown()
	}
	m.ScrollDown()
	if m.wheel.State() != minSettingValue {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", minSettingValue, m.wheel.State())
	}
}

func TestMouse_Click(t *testing.T) {
	s := NewScreen(Width, Height)
	m := NewMouse(*s, &Button{}, &Button{}, NewWheel())
	m.LeftButtonDown()
	if m.leftButton.State() != true {
		t.Errorf("Error button click. Expected - %v, got - %v", true, m.leftButton.State())
	}
	m.LeftButtonUp()
	if m.leftButton.State() != false {
		t.Errorf("Error button click. Expected - %v, got - %v", false, m.leftButton.State())
	}
	m.RightButtonDown()
	if m.rightButton.State() != true {
		t.Errorf("Error button click. Expected - %v, got - %v", true, m.rightButton.State())
	}
	m.RightButtonUp()
	if m.rightButton.State() != false {
		t.Errorf("Error button click. Expected - %v, got - %v", false, m.rightButton.State())
	}
}

func TestMouse_Info(t *testing.T) {
	s := NewScreen(Width, Height)
	m := NewMouse(*s, &Button{}, &Button{}, NewWheel())
	var err error
	m.Sensitivity(5)
	m.LeftButtonDown()
	for i := 0; i < 4; i++ {
		m.ScrollUp()
	}

	oldOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.Info()

	err = w.Close()
	if err != nil {
		t.Errorf("Error closing pipe: %s\n", err)
	}

	os.Stdout = oldOutput

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Errorf("Error while copy: %s\n", err)
	}

	if equal := strings.Compare(testInfoStr, buf.String()); equal != 0 {
		t.Errorf("Error info output. Expected:\n%s\nGot:\n%s", testInfoStr, buf.String())
	}
}

func TestMouse_Reset(t *testing.T) {
	s := NewScreen(Width, Height)
	m := NewMouse(*s, &Button{}, &Button{}, NewWheel())
	var err error
	m.Move(1, 1, *s)
	m.Sensitivity(4)
	m.RightButtonDown()
	m.ScrollUp()
	m.ScrollUp()
	m.Reset(*s)
	var str = `Mouse information:
X position - 512
Y position - 384
Sensitivity - 1
Is left button pressed? - false
Is right button pressed? - false
Scroll value - 5
`
	oldOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.Info()

	err = w.Close()
	if err != nil {
		t.Errorf("Error closing pipe: %s\n", err)
	}
	os.Stdout = oldOutput

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Errorf("Error while copy: %s\n", err)
	}

	if equal := strings.Compare(str, buf.String()); equal != 0 {
		t.Errorf("Error info output. Expected:\n%s\nGot:\n%s", str, buf.String())
	}
}
