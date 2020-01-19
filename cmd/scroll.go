package cmd

import (
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type mouseWheelScroller struct {
	mouse *mouse.Mouse
}

func NewMouseWheelScroller(mouse *mouse.Mouse) *cobra.Command {
	mouseWheelScroller := &mouseWheelScroller{mouse}

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
	result, err := selectScrollDirection()
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

func selectScrollDirection() (result string, err error) {
	prompt := promptui.Select{
		Label: "Select scroll direction",
		Items: []string{"Up", "Down"},
	}

	_, result, err = prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
