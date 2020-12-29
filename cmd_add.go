package kuronekocat

import "github.com/spf13/cobra"

func newAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add --id 4233-999-888 --info computer",
		Short: "add id, info",

		RunE: func(_ *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
