package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// sensCmd represents the sens command
var sensCmd = &cobra.Command{
	Use:   "sens",
	Short: "set sensitivity of mouse",
	Long: `Sens command set mouse sensitivity to value that are entered from the keyboard
`,
	RunE: func(cmd *cobra.Command, args []string) error{
		var sensStr string
		fmt.Print("Enter sensitivity: ")
		fmt.Scan(&sensStr)
		sens, err := strconv.ParseUint(sensStr, 10, 8)
		if err != nil {
			return fmt.Errorf("invalid sensitivity: %s", err)
		}
		m.SetSensitivity(uint8(sens))
		return nil
	},
}
