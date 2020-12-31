package kuronekocat

import (
	"github.com/spf13/cobra"
)

const cmdName = "kuronekocat"

func NewKuronekoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  cmdName,
		Long: "Kuroneko Yamato Tracking Wrapper Tool",
	}

	cmd.AddCommand(
		newGetCmd(),
		newAddCmd(),
		newQueueCmd(),
		newReceiveCmd(),
	)

	return cmd
}
