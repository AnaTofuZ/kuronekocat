package kuronekocat

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List current ordered queues",

		RunE: execListCmd,
	}
	return cmd
}

func execListCmd(_ *cobra.Command, arg []string) error {
	orders, err := readFromHomeJSON()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"品物", "追跡番号"})

	for _, order := range orders {
		table.Append([]string{order.Explain, order.ID})
	}
	table.Render()
	return nil
}
