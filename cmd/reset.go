package cmd

import (
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "restore default settings",
	Long: `Reset command restores default settings
`,
	Run: func(cmd *cobra.Command, args []string) {
		m.Reset(*s)
	},
}

