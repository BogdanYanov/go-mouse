package main

import (
	"fmt"
	"github.com/BogdanYanov/go-mouse/cmd"
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var (
		rootCmd = &cobra.Command{
			Use: "go-mouse",
			Long: `TestCLI is a CLI program, which simulates the behavior of the mouse
`,
		}
		prompter *cmd.Prompter
	)

	screenCmd := mouse.NewScreen(mouse.Width, mouse.Height)
	mouseCmd := mouse.NewMouse(*screenCmd)

	rootCmd.AddCommand(
		cmd.NewButtonPresser(mouseCmd, prompter),
		cmd.NewButtonReleaser(mouseCmd, prompter),
		cmd.NewMouseMover(mouseCmd, screenCmd),
		cmd.NewMouseInformer(mouseCmd),
		cmd.NewMouseRestorer(mouseCmd, screenCmd),
		cmd.NewMouseWheelScroller(mouseCmd, prompter),
		cmd.NewMouseSensor(mouseCmd))

	prompter = cmd.NewPrompter(rootCmd)

	for {
		if err := prompter.SelectMenu(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if err := rootCmd.Execute(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}
