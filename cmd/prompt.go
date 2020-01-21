package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

type Prompter struct {
	command *cobra.Command
}

func NewPrompter(command *cobra.Command) *Prompter {
	return &Prompter{command}
}

func (p *Prompter) getCommands() (commandsNames []string) {
	commands := p.command.Commands()
	commandsNames = make([]string, 0, len(commands)+2)
	for _, cmd := range commands {
		commandsNames = append(commandsNames, cmd.Name())
	}
	return commandsNames
}

func (p *Prompter) parseResult(result string) {
	if result == "exit" {
		os.Exit(0)
	}
	if result == "about" {
		result = "help"
	}
	if len(os.Args) > 1 {
		os.Args = os.Args[:2]
		os.Args[1] = result
	} else {
		os.Args = append(os.Args, result)
	}
}

func (p *Prompter) SelectMenu() error {
	commands := p.getCommands()
	commands = append(commands, "about", "exit")

	prompt := promptui.Select{
		Label: "Select Command",
		Items: commands,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("Prompt failed %v\n", err)
	}

	err = clearConsole()
	if err != nil {
		return err
	}

	p.parseResult(result)

	return nil
}

func (p *Prompter) selectBtn() (result string, err error) {
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

func (p *Prompter) selectScrollDirection() (result string, err error) {
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

func clearConsole() error {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
