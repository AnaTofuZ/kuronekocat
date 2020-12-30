package kuronekocat

import "github.com/spf13/cobra"

func newQueueCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "queue",
		Short: "get kuroneko infomations",

		RunE: execQueueCmd,
	}
	return cmd
}

func execQueueCmd(_ *cobra.Command, arg []string) error {
	_, err := readFromHomeJSON()
	if err != nil {
		return err
	}
	return nil
}
