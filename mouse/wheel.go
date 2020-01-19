package mouse

// Wheel contains scroll value of mouse wheel.
type Wheel struct {
	scrollValue uint8
}

func (w *Wheel) scrollUp() {
	if w.scrollValue != maxSettingValue {
		w.scrollValue++
	} else {
		w.scrollValue = maxSettingValue
	}
}

func (w *Wheel) scrollDown() {
	if w.scrollValue != minSettingValue {
		w.scrollValue--
	} else {
		w.scrollValue = minSettingValue
	}
}
