package cmd

import (
	"github.com/ILLIDOM/netutils/cmd/multicast"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "netutils",
		Short: "CLI",
		Long:  "netutils for network engineers",
	}

	rootCmd.AddCommand(
		multicast.New(),
	)
}

func Execute() error {
	return New().Execute()
}
