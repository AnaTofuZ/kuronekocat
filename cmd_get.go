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
	var explains []string
	if len(orderNumbers) == 0 {
		orders, err := readFromHomeJSON()
		if err != nil {
			return err
		}
		for _, ord := range orders {
			orderNumbers = append(orderNumbers, ord.ID)
			explains = append(explains, ord.Explain)
		}
	}

	fields, err := getFromTneko(orderNumbers)
	if err != nil {
		return err
	}
	if len(explains) == 0 {
		showTable(fields)
	} else {
		showTableWExplain(fields, explains)
	}
	return nil
}

func showTable(infos *parsedInfoType) {
	header := infos.headers
	fields := infos.orders

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header[:])
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetRowLine(true)
	for _, order := range fields {
		for _, v := range order[:] {
			table.Append(v[:])
		}
	}
	table.Render()
}

func showTableWExplain(infos *parsedInfoType, explains []string) {
	header := infos.headers[:]
	fields := infos.orders

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(append([]string{"品物"}, header...))
	table.SetAutoMergeCellsByColumnIndex([]int{0, 1})
	table.SetRowLine(true)
	for i, order := range fields {
		for _, v := range order[:] {
			feeld := v[:]
			table.Append(append([]string{explains[i]}, feeld...))
		}
	}
	table.Render()
}
