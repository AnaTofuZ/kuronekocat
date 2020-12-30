package kuronekocat

import (
	"github.com/spf13/cobra"
)

const cmdName = "kuronekocat"

var rootCmd = &cobra.Command{
	Use:  cmdName,
	Long: "クロネコヤマトから情報をとるくん",
}

func NewKuronekoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "kuronekocat",
	}

	cmd.AddCommand(
		newGetCmd(),
		newAddCmd(),
	)

	return cmd
}
