package cmd

import (
	"bufio"
	"fmt"
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

type mouseMover struct {
	mouse  *mouse.Mouse
	screen *mouse.Screen
}

// NewMouseMover create a new cobra command that moves the mouse cursor to the coordinates specified from the console.
func NewMouseMover(mouse *mouse.Mouse, screen *mouse.Screen) *cobra.Command {
	mouseMover := &mouseMover{mouse, screen}

	var moveCmd = &cobra.Command{
		Use:   "move",
		Short: "changes x y mouse coordinates",
		Long: `Move command changes x y coordinates of mouse to values that are entered from the keyboard
`,
		RunE: mouseMover.Exec,
	}
	return moveCmd
}

func (mm *mouseMover) Exec(cmd *cobra.Command, args []string) error {
	var (
		consoleStr  string
		console     = bufio.NewReader(os.Stdin)
		coordinates = make([]string, 0, 2)
	)

	fmt.Print("Enter x, y coords to moving, using space: ")
	consoleStr, err := console.ReadString('\n')
	if err != nil {
		return err
	}

	consoleStr = strings.Replace(consoleStr, "\n", "", -1)
	coordinates = strings.Split(consoleStr, " ")

	x, err := strconv.ParseUint(coordinates[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid coordinates: %s", err)
	}

	y, err := strconv.ParseUint(coordinates[1], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid coordinates: %s", err)
	}

	mm.mouse.Move(uint32(x), uint32(y), *mm.screen)

	return nil
}
