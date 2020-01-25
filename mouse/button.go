package mouse

// Button implements Up, Down and contains a button pressed state
type button struct {
	isPressed bool
}

func (btn *button) down() {
	btn.isPressed = true
}

func (btn *button) up() {
	btn.isPressed = false
}
