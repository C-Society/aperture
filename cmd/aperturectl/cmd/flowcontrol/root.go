package flowcontrol

import (
	"github.com/spf13/cobra"

	"github.com/fluxninja/aperture/cmd/aperturectl/cmd/utils"
)

var controller utils.ControllerConn

func init() {
	controller.InitFlags(FlowControlCmd.PersistentFlags())

	FlowControlCmd.AddCommand(ListCmd)
	FlowControlCmd.AddCommand(PreviewCmd)
}

// FlowControlCmd is the command to observe Flow Control control points.
var FlowControlCmd = &cobra.Command{
	Use:               "flow-control",
	Short:             "Flow Control integrations",
	Long:              `Use this command to query information about active Flow Control integrations`,
	SilenceErrors:     true,
	PersistentPreRunE: controller.PreRunE,
	PersistentPostRun: controller.PostRun,
}
