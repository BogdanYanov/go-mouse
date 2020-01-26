package mouse

type Clicker interface {
	Up()
	Down()
	State() bool
}

// Button implements Up, Down and contains a button pressed state
type Button struct {
	IsPressed bool
}

func (btn *Button) Down() {
	btn.IsPressed = true
}

func (btn *Button) Up() {
	btn.IsPressed = false
}

func (btn *Button) State() bool {
	return btn.IsPressed
}