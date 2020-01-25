package cmd

import (
	"bufio"
	"fmt"
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

type mouseSensor struct {
	mouse *mouse.Mouse
}

// NewMouseSensor create a new cobra command which sets the mouse sensitivity value according to the value entered from the console.
func NewMouseSensor(mouse *mouse.Mouse) *cobra.Command {
	mouseSensor := &mouseSensor{mouse}

	var sensCmd = &cobra.Command{
		Use:   "sens",
		Short: "set sensitivity of mouse",
		Long: `Sens command set mouse sensitivity to value that are entered from the keyboard
`,
		RunE: mouseSensor.Exec,
	}

	return sensCmd
}

func (ms *mouseSensor) Exec(cmd *cobra.Command, args []string) error {
	var (
		sensStr string
		console = bufio.NewReader(os.Stdin)
	)

	fmt.Print("Enter sensitivity: ")
	sensStr, err := console.ReadString('\n')
	if err != nil {
		return err
	}

	sensStr = strings.Replace(sensStr, "\n", "", -1)

	sens, err := strconv.ParseUint(sensStr, 10, 8)
	if err != nil {
		return fmt.Errorf("invalid sensitivity: %s", err)
	}

	ms.mouse.Sensitivity(uint8(sens))

	return nil
}
