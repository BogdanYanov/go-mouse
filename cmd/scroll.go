package cmd

import (
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/spf13/cobra"
)

type mouseWheelScroller struct {
	mouse *mouse.Mouse
	prompter *Prompter
}

func NewMouseWheelScroller(mouse *mouse.Mouse, prompter *Prompter) *cobra.Command {
	mouseWheelScroller := &mouseWheelScroller{mouse, prompter}

	var scrollCmd = &cobra.Command{
		Use:   "scroll",
		Short: "scrolls the mouse wheel",
		Long: `Scroll command makes a simulation of a mouse wheel scrolls
`,
		RunE: mouseWheelScroller.Exec,
	}

	return scrollCmd
}

func (mws *mouseWheelScroller) Exec(cmd *cobra.Command, args []string) error {
	result, err := mws.prompter.selectScrollDirection()
	if err != nil {
		return err
	}

	switch result {
	case "Up":
		mws.mouse.ScrollUp()
	case "Down":
		mws.mouse.ScrollDown()
	}

	return nil
}

