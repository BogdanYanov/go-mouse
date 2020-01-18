package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// scrollCmd represents the scroll command
var scrollCmd = &cobra.Command{
	Use:   "scroll",
	Short: "scrolls the mouse wheel",
	Long: `Scroll command makes a simulation of a mouse wheel scrolls
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := selectScrollDirection()
		if err != nil {
			return err
		}
		switch result {
		case "Up":
			m.ScrollUp()
		case "Down":
			m.ScrollDown()
		}
		return nil
	},
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
