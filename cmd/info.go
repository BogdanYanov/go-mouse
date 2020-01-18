package cmd

import (
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show information about mouse settings and states",
	Long: `Info command show information about mouse settings and states
`,
	Run: func(cmd *cobra.Command, args []string) {
		m.Info()
	},
}
