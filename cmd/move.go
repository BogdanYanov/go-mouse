package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "changes x y mouse coordinates",
	Long: `Move command changes x y coordinates of mouse to values that are entered from the keyboard
`,
	RunE: func(cmd *cobra.Command, args []string) error{
		var xStr, yStr string
		fmt.Print("Enter x, y coords to moving, using space: ")
		fmt.Scan(&xStr, &yStr)
		x, err := strconv.ParseUint(xStr, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid coordinates: %s", err)
		}
		y, err := strconv.ParseUint(yStr, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid coordinates: %s", err)
		}
		m.Move(uint32(x), uint32(y), *s)
		return nil
	},
}


