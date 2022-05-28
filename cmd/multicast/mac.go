package multicast

import (
	"fmt"
	"net"
	"os"

	"github.com/ILLIDOM/netutils/utils"

	"github.com/spf13/cobra"
)

var multicastPrefixMAC = []byte{1, 0, 94} //Decimal Value of 01, 00, 5E

func newMacCommand() *cobra.Command {
	macCommand := &cobra.Command{
		Use:   "mac",
		Short: "mac multicast",
		Example: `provide single mac address (Ethertype 0x0800) to get all corresponding IP multicast addresses
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

			if !isValidMulticastMAC(hwAddress) {
				fmt.Println("No valid Multicast MAC provided: remember Multicast MAC addresses start with 01:00:5e")
				return
			}

			utils.MulticastIPfromMAC(hwAddress)
		},
	}
	return macCommand
}

func isValidMulticastMAC(macAddress net.HardwareAddr) bool {
	for i, v := range multicastPrefixMAC {
		if v != macAddress[i] {
			return false
		}
	}
	return true
}
