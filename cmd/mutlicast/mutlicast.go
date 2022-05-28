package multicast

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	multicastCmd := &cobra.Command{
		Use:   "multicast",
		Short: "multicast informations",
	}
	multicastCmd.AddCommand(
		newMacCommand(),
	)
	return multicastCmd
}
