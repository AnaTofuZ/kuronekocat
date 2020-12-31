package kuronekocat

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func newQueueCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "queue",
		Short: "今溜まっている品物一覧を表示",

		RunE: execQueueCmd,
	}
	return cmd
}

func execQueueCmd(_ *cobra.Command, arg []string) error {
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
