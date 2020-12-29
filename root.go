package kuronekocat

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "kuronekocat",
	Long: "クロネコヤマトから情報をとるくん",
}

func NewKuronekoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "kuronekocat",
	}

	cmd.AddCommand(
		newGetCmd(),
	)

	return cmd
}
