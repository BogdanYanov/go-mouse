package mouse

// Clicker is the interface that wraps mouse buttons functions Up and Down.
type Clicker interface {
	up()
	down()
}

// Button implements Up, Down and contains a button pressed state
type Button struct {
	btnPressed bool
}

// Down simulates a button press
func (btn *Button) down() {
	btn.btnPressed = true
}

// Up simulates a button release
func (btn *Button) up() {
	btn.btnPressed = false
}

