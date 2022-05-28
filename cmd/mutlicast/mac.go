package multicast

import "github.com/spf13/cobra"

func newMacCommand() *cobra.Command {
	macCommand := &cobra.Command{
		Use:   "mac",
		Short: "mac multicast",
	}
	return macCommand
}
