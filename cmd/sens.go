package cmd

import (
	"fmt"
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/spf13/cobra"
	"strconv"
)

type mouseSensor struct {
	mouse *mouse.Mouse
}

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
	var sensStr string

	fmt.Print("Enter sensitivity: ")
	_, err := fmt.Scan(&sensStr)
	if err != nil {
		return err
	}

	sens, err := strconv.ParseUint(sensStr, 10, 8)
	if err != nil {
		return fmt.Errorf("invalid sensitivity: %s", err)
	}

	ms.mouse.SetSensitivity(uint8(sens))

	return nil
}
