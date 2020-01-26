package mouse

type Scroller interface {
	ScrollUp()
	ScrollDown()
	State() uint8
}

type Wheel struct {
	ScrollValue uint8
}

func NewWheel() *Wheel {
	return &Wheel{defaultSettingValue}
}

func (w *Wheel) ScrollUp() {
	if w.ScrollValue != maxSettingValue {
		w.ScrollValue++
	} else {
		w.ScrollValue = maxSettingValue
	}
}

func (w *Wheel) ScrollDown() {
	if w.ScrollValue != minSettingValue {
		w.ScrollValue--
	} else {
		w.ScrollValue = minSettingValue
	}
}

func (w *Wheel) State() uint8{
	return w.ScrollValue
}

