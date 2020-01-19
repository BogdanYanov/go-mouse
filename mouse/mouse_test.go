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
	s := NewScreen()
	m := NewMouse(*s)
	if m == nil {
		t.Errorf("mouse is nil")
	}
}

func TestMouse_Move(t *testing.T) {
	s := NewScreen()
	m := NewMouse(*s)
	m.Move(1260, 1024, *s)
	if m.posX != s.width || m.posY != s.height {
		t.Errorf("Coords out of screen. Expected X - %d, Y - %d; Get X - %d, Y - %d",
			s.width,
			s.height,
			m.posX,
			m.posY)
	}
	m.SetSensitivity(10)
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
	s := NewScreen()
	m := NewMouse(*s)
	m.SetSensitivity(11)
	if m.sensitivity != 10 {
		t.Errorf("Sensitivity out of range. Expected sensitivity - %d, get - %d", 10, m.sensitivity)
	}
	m.SetSensitivity(0)
	if m.sensitivity != 1 {
		t.Errorf("Sensitivity out of range. Expected sensitivity - %d, get - %d", 1, m.sensitivity)
	}
}

func TestMouse_WheelScrollUp(t *testing.T) {
	s := NewScreen()
	m := NewMouse(*s)
	m.wheel.scrollValue = 7
	m.ScrollUp()
	if m.wheel.scrollValue != 8 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", 8, m.wheel.scrollValue)
	}
	for i := 0; i < 4; i++ {
		m.ScrollUp()
	}
	if m.wheel.scrollValue != 10 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", 10, m.wheel.scrollValue)
	}
}

func TestMouse_WheelScrollDown(t *testing.T) {
	s := NewScreen()
	m := NewMouse(*s)
	m.wheel.scrollValue = 3
	m.ScrollDown()
	if m.wheel.scrollValue != 2 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", 2, m.wheel.scrollValue)
	}
	m.ScrollDown()
	m.ScrollDown()
	if m.wheel.scrollValue != 1 {
		t.Errorf("Error scrolling. Scroll value expected - %d, get - %d", 1, m.wheel.scrollValue)
	}
}

func TestMouse_Click(t *testing.T) {
	s := NewScreen()
	m := NewMouse(*s)
	m.LeftBtnDown()
	if m.leftBtn.btnPressed != true {
		t.Errorf("Error button click. Expected - %v, got - %v", true, m.leftBtn.btnPressed)
	}
	m.LeftBtnUp()
	if m.leftBtn.btnPressed != false {
		t.Errorf("Error button click. Expected - %v, got - %v", false, m.leftBtn.btnPressed)
	}
	m.RightBtnDown()
	if m.rightBtn.btnPressed != true {
		t.Errorf("Error button click. Expected - %v, got - %v", true, m.rightBtn.btnPressed)
	}
	m.RightBtnUp()
	if m.rightBtn.btnPressed != false {
		t.Errorf("Error button click. Expected - %v, got - %v", false, m.rightBtn.btnPressed)
	}
}

func TestMouse_Info(t *testing.T) {
	s := NewScreen()
	m := NewMouse(*s)
	var err error
	m.SetSensitivity(5)
	m.LeftBtnDown()
	m.wheel.scrollValue = 8
	m.ScrollUp()

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
	s := NewScreen()
	m := NewMouse(*s)
	var err error
	m.Move(1, 1, *s)
	m.SetSensitivity(4)
	m.RightBtnDown()
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
