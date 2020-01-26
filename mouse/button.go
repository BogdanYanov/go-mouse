package mouse

// Clicker is the interface which contains required methods to use as a button.
//
// The structure that implement this interface must store the boolean state of the button click.
type Clicker interface {
	Up()
	Down()
	State() bool
}

// Button implements Clicker interface methods and contains a button pressed state.
type Button struct {
	IsPressed bool
}

// Down simulates the button pressed and changed button pressed state to true.
func (btn *Button) Down() {
	btn.IsPressed = true
}

// Up simulates the button pushed and changed button pressed state to false.
func (btn *Button) Up() {
	btn.IsPressed = false
}

// State return button pressed state.
func (btn *Button) State() bool {
	return btn.IsPressed
}
