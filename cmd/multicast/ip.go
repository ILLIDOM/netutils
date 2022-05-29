package multicast

import (
	"fmt"
	"net"
	"os"

	"github.com/ILLIDOM/netutils/utils"
	"github.com/spf13/cobra"
)

func newIpCommand() *cobra.Command {
	ipCommand := &cobra.Command{
		Use:   "ip",
		Short: "ip multicast",
		Example: `provide a single ipv4 address to get the corresponding MAC multicast addresses
		e.g. 224.0.0.1 (case insensitive)
		`,
		Run: func(cmd *cobra.Command, args []string) {
			// check if input is empty
			if len(args) != 1 {
				fmt.Println("Please provide a single MAC address")
				return
			}

			ipAddressString := args[0]
			ipAddress := net.ParseIP(ipAddressString)
			if ipAddress == nil {
				fmt.Println("Error: can't parse IP address")
				os.Exit(1)
			}

			if !utils.IsValidMulticastIP(ipAddress) {
				fmt.Println("No valid multicast IP provided: remember the multicast range is 224.0.0.0 – 239.255.255.255")
				return
			}

			multicastMACAddress := utils.MulticastMACFromIP(ipAddress)
			if all {
				fmt.Printf("Multicast IP addresses mapping to %v\n", multicastMACAddress)
				allMulticastIPs := utils.MulticastIPfromMAC(multicastMACAddress)
				for _, v := range allMulticastIPs {
					fmt.Println(utils.BinaryIpStringToIntString(v))
				}
			} else {
				fmt.Println(multicastMACAddress)
			}
		},
	}
	ipCommand.Flags().BoolVar(&all, "all", false, "Flag to show all 32 mutlicast IPs")
	return ipCommand
}
