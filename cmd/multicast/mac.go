package multicast

import (
	"fmt"
	"net"
	"os"

	"github.com/ILLIDOM/netutils/utils"

	"github.com/spf13/cobra"
)

var all bool //flag indicating if all 32 mc ips are printed

func newMacCommand() *cobra.Command {
	macCommand := &cobra.Command{
		Use:   "mac",
		Short: "mac multicast",
		Example: `provide a single mac address (Ethertype 0x0800) to get all corresponding IP multicast addresses
		e.g. 01:00:5E:DD:EE:FF or 01-00-5E-DD-EE-FF (case insensitive)
		`,

		Run: func(cmd *cobra.Command, args []string) {
			// check if input is empty
			if len(args) != 1 {
				fmt.Println("Please provide a single MAC address")
				return
			}

			macAddress := args[0]

			hwAddress, err := net.ParseMAC(macAddress)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			if !utils.IsValidMulticastMAC(hwAddress) {
				fmt.Println("No valid Multicast MAC provided: remember Multicast MAC addresses start with 01:00:5e")
				return
			}

			allMulticastIPs := utils.MulticastIPfromMAC(hwAddress)

			if all {
				for _, v := range allMulticastIPs {
					fmt.Println(utils.BinaryIpStringToIntString(v))
				}
			} else {
				fmt.Println(utils.BinaryIpStringToIntString(allMulticastIPs[0]))
			}
		},
	}
	macCommand.Flags().BoolVar(&all, "all", false, "Flag to show all 32 mutlicast IPs")
	return macCommand
}
