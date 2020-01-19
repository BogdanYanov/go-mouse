package mouse

// Clicker is the interface that wraps mouse buttons functions Up and Down.
type Clicker interface {
	up()
	down()
}

// Scroller is the interface that wraps ScrollUp and ScrollDown methods
type Scroller interface {
	scrollUp()
	scrollDown()
}
