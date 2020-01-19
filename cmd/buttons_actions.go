package cmd

import (
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type buttonReleaser struct {
	mouse *mouse.Mouse
}

type buttonPresser struct {
	mouse *mouse.Mouse
}

func selectBtn() (result string, err error) {
	prompt := promptui.Select{
		Label: "Select button",
		Items: []string{"Right", "Left", "Both"},
	}

	_, result, err = prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func NewButtonReleaser(mouse *mouse.Mouse) *cobra.Command {
	buttonReleaser := &buttonReleaser{mouse}

	var btnUpCmd = &cobra.Command{
		Use:   "btn-up",
		Short: "release the mouse button(s)",
		Long: `Btn-up command makes a simulation of a mouse button(s) release
`,
		RunE: buttonReleaser.Exec,
	}

	return btnUpCmd
}

func NewButtonPresser(mouse *mouse.Mouse) *cobra.Command {
	buttonPresser := &buttonPresser{mouse}

	btnDownCmd := &cobra.Command{
		Use:   "btn-down",
		Short: "presses a mouse button(s)",
		Long: `Btn-down command makes a simulation of mouse button(s) press
`,
		RunE: buttonPresser.Exec,
	}

	return btnDownCmd
}

func (br *buttonReleaser) Exec(cmd *cobra.Command, args []string) error {
	result, err := selectBtn()
	if err != nil {
		return err
	}

	switch result {
	case "Right":
		br.mouse.RightBtnUp()
	case "Left":
		br.mouse.LeftBtnUp()
	case "Both":
		br.mouse.RightBtnUp()
		br.mouse.LeftBtnUp()
	}

	return nil
}

func (bp *buttonPresser) Exec(cmd *cobra.Command, args []string) error {
	result, err := selectBtn()
	if err != nil {
		return err
	}

	switch result {
	case "Right":
		bp.mouse.RightBtnDown()
	case "Left":
		bp.mouse.LeftBtnDown()
	case "Both":
		bp.mouse.RightBtnDown()
		bp.mouse.LeftBtnDown()
	}

	return nil
}
