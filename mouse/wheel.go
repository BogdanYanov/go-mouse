package mouse

// Scroller is the interface which contains required methods to use as a mouse wheel.
//
// The structure that implement this interface must store the mouse wheel scroll value.
type Scroller interface {
	ScrollUp()
	ScrollDown()
	State() uint8
}

// Wheel implements Scroller interface methods and contains a scroll value.
type Wheel struct {
	ScrollValue uint8
}

// NewWheel returns new Wheel structure object with a scroll value by default.
func NewWheel() *Wheel {
	return &Wheel{defaultSettingValue}
}

// ScrollUp simulates mouse wheel scroll up.
func (w *Wheel) ScrollUp() {
	if w.ScrollValue != maxSettingValue {
		w.ScrollValue++
	} else {
		w.ScrollValue = maxSettingValue
	}
}

// ScrollDown simulates mouse wheel scroll down.
func (w *Wheel) ScrollDown() {
	if w.ScrollValue != minSettingValue {
		w.ScrollValue--
	} else {
		w.ScrollValue = minSettingValue
	}
}

// State return mouse wheel scroll value
func (w *Wheel) State() uint8 {
	return w.ScrollValue
}
