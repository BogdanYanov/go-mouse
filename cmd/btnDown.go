package cmd

import (
	"github.com/spf13/cobra"
)

// btnDownCmd represents the btnDown command
var btnDownCmd = &cobra.Command{
	Use:   "btn-down",
	Short: "presses a mouse button(s)",
	Long: `Btn-down command makes a simulation of mouse button(s) press
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := selectBtn()
		if err != nil {
			return err
		}
		switch result {
		case "Right":
			m.RightBtnDown()
		case "Left":
			m.LeftBtnDown()
		case "Both":
			m.RightBtnDown()
			m.LeftBtnDown()
		}
		return nil
	},
}
