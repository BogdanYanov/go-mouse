package cmd

import (
	"github.com/BogdanYanov/go-mouse/mouse"
	"github.com/spf13/cobra"
)

type mouseInformer struct {
	mouse *mouse.Mouse
}

func NewMouseInformer(mouse *mouse.Mouse) *cobra.Command {
	mouseInformer := &mouseInformer{mouse}

	var infoCmd = &cobra.Command{
		Use:   "info",
		Short: "show information about mouse settings and states",
		Long: `Info command show information about mouse settings and states
`,
		Run: mouseInformer.Exec,
	}

	return infoCmd
}

func (mi *mouseInformer) Exec(cmd *cobra.Command, args []string) {
	mi.mouse.Info()
}
