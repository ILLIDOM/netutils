package cmd

import (
	"github.com/ILLIDOM/netutils/multicast"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "neturils",
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
