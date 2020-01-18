package cmd

import (
	"fmt"
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	width uint32 = 1024
	height uint32 = 768
)

var s *mouse.Screen
var m *mouse.Mouse

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "testCLI",
	Short: "Mouse CLI program",
	Long: `TestCLI is a CLI program, which simulates the behavior of the mouse
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	s = mouse.NewScreen(width, height)
	m = mouse.NewMouse(*s)
	for{
		result := runSelectMenu()
		parseResult(result)
		if err := rootCmd.Execute(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mycli.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(moveCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(btnUpCmd)
	rootCmd.AddCommand(btnDownCmd)
	rootCmd.AddCommand(scrollCmd)
	rootCmd.AddCommand(sensCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mycli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mycli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func runSelectMenu() (result string) {
	prompt := promptui.Select{
		Label: "Select Command",
		Items: []string{"move", "info", "reset", "btn-up", "btn-down", "scroll", "sens", "about", "exit"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}
	return result
}

func parseResult(result string) {
	if result == "exit" {
		os.Exit(0)
	}
	if result == "about" {
		result = "help"
	}
	if len(os.Args) == 1 {
		os.Args = append(os.Args, result)
	} else {
		os.Args = os.Args[:2]
		os.Args[1] = result
	}
}