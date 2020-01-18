package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// btnUpCmd represents the btnUp command
var btnUpCmd = &cobra.Command{
	Use:   "btn-up",
	Short: "release the mouse button(s)",
	Long: `Btn-up command makes a simulation of a mouse button(s) release
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := selectBtn()
		if err != nil {
			return err
		}
		switch result {
		case "Right":
			m.RightBtnUp()
		case "Left":
			m.LeftBtnUp()
		case "Both":
			m.RightBtnUp()
			m.LeftBtnUp()
		}
		return nil
	},
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
