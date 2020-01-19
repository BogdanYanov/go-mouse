package cmd

import (
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/spf13/cobra"
)

type mouseRestorer struct {
	mouse  *mouse.Mouse
	screen *mouse.Screen
}

func NewMouseRestorer(mouse *mouse.Mouse, screen *mouse.Screen) *cobra.Command {
	mouseRestorer := &mouseRestorer{mouse, screen}

	var resetCmd = &cobra.Command{
		Use:   "reset",
		Short: "restore default settings",
		Long: `Reset command restores default settings
`,
		Run: mouseRestorer.Exec,
	}

	return resetCmd
}

func (mr *mouseRestorer) Exec(cmd *cobra.Command, args []string) {
	mr.mouse.Reset(*mr.screen)
}
