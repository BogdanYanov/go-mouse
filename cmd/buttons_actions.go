package cmd

import (
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/spf13/cobra"
)

type buttonReleaser struct {
	mouse *mouse.Mouse
	prompter *Prompter
}

type buttonPresser struct {
	mouse *mouse.Mouse
	prompter *Prompter
}

func NewButtonReleaser(mouse *mouse.Mouse, prompter *Prompter) *cobra.Command {
	buttonReleaser := &buttonReleaser{mouse, prompter}

	var btnUpCmd = &cobra.Command{
		Use:   "btn-up",
		Short: "release the mouse button(s)",
		Long: `Btn-up command makes a simulation of a mouse button(s) release
`,
		RunE: buttonReleaser.Exec,
	}

	return btnUpCmd
}

func NewButtonPresser(mouse *mouse.Mouse, prompter *Prompter) *cobra.Command {
	buttonPresser := &buttonPresser{mouse, prompter}

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
	result, err := br.prompter.selectBtn()
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
	result, err := bp.prompter.selectBtn()
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
