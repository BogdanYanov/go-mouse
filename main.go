package main

import (
	"fmt"
	"github.com/BogdanYanov/go-mouse/cmd"
	. "github.com/BogdanYanov/go-mouse/mouse"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use: "go-mouse",
		Long: `TestCLI is a CLI program, which simulates the behavior of the mouse
`,
	}
	prompter *cmd.Prompter
)

func init() {
	screen := NewScreen()
	mouse := NewMouse(*screen)
	rootCmd.AddCommand(
		cmd.NewButtonPresser(mouse),
		cmd.NewButtonReleaser(mouse),
		cmd.NewMouseMover(mouse, screen),
		cmd.NewMouseInformer(mouse),
		cmd.NewMouseRestorer(mouse, screen),
		cmd.NewMouseWheelScroller(mouse),
		cmd.NewMouseSensor(mouse))
	prompter = cmd.NewPrompter(rootCmd)
}

func main() {
	for {
		err := prompter.SelectMenu()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if err := rootCmd.Execute(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}
