package cmd

import (
	"fmt"
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/spf13/cobra"
	"strconv"
)

type mouseMover struct {
	mouse  *mouse.Mouse
	screen *mouse.Screen
}

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
	var xStr, yStr string

	fmt.Print("Enter x, y coords to moving, using space: ")
	_, err := fmt.Scan(&xStr, &yStr)
	if err != nil {
		return err
	}

	x, err := strconv.ParseUint(xStr, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid coordinates: %s", err)
	}

	y, err := strconv.ParseUint(yStr, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid coordinates: %s", err)
	}

	mm.mouse.Move(uint32(x), uint32(y), *mm.screen)

	return nil
}
