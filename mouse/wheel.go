package mouse

type wheel struct {
	scrollValue uint8
}

func (w *wheel) scrollUp() {
	if w.scrollValue != maxSettingValue {
		w.scrollValue++
	} else {
		w.scrollValue = maxSettingValue
	}
}

func (w *wheel) scrollDown() {
	if w.scrollValue != minSettingValue {
		w.scrollValue--
	} else {
		w.scrollValue = minSettingValue
	}
}
