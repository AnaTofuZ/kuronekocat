package kuronekocat

import (
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

var addQuery = order{}

func newAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add --number 4233-999-888 --info computer",
		Short: "add id, info",

		RunE: addCmd,
	}

	cmd.Flags().StringVar(&addQuery.ID, "number", "", "伝票番号")
	cmd.Flags().StringVar(&addQuery.Explain, "explain", "", "説明")
	return cmd
}

func addCmd(_ *cobra.Command, _ []string) error {
	if addQuery.ID == "" {
		return xerrors.New("[error] require --number")
	}

	if addQuery.Explain == "" {
		addQuery.Explain = addQuery.ID
	}

	addQuery.ID = removeHyphen(addQuery.ID)
	if err := addOrderToJSON(addQuery); err != nil {
		return xerrors.Errorf("[error] %+w", err)
	}

	return nil
}

func removeHyphen(id string) string {
	return strings.ReplaceAll(id, "-", "")
}
