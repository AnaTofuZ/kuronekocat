package kuronekocat

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [ids,...]",
		Short: "get kuroneko infomations",

		RunE: execGetCmd,
	}
	return cmd
}

func execGetCmd(_ *cobra.Command, arg []string) error {
	orderNumbers := arg
	if len(orderNumbers) == 0 {
		orderNumbers = []string{"11111"}
	}

	fields, err := getFromTneko(orderNumbers)
	if err != nil {
		return err
	}
	showTable(fields)
	return nil
}

func showTable(infos *parsedInfoType) {
	header := infos.headers
	fields := infos.orders

	for _, order := range fields {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(header[:])
		for _, v := range order[:] {
			table.Append(v[:])
		}
		table.Render()
	}
}
