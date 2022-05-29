package utils

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

const multicastIpPrefix = "1110"

var multicastPrefixMAC = []byte{1, 0, 94} //Decimal Value of 01, 00, 5E

func IsValidMulticastMAC(macAddress net.HardwareAddr) bool {
	// check first 3 bytes -> need to match multicast prefix
	for i, v := range multicastPrefixMAC {
		if v != macAddress[i] {
			return false
		}
	}
	return true
}

func IsValidMulticastIP(ipAddress net.IP) bool {
	octet1 := ipAddress[12]         //net.IP is always 16 Bytes long - last 4 bytes are ipv4 address
	octet1Shifted := octet1 >> 4    // shift octet to get 0000 + last 4 bits
	prefixShifted := byte(224) >> 4 // shit to get prefix -> 00001110

	return octet1Shifted == prefixShifted
}

func MulticastMACFromIP(ipAddress net.IP) net.HardwareAddr {
	var multicastMAC = []byte{}
	multicastMAC = append(multicastMAC, multicastPrefixMAC...) //add fix first 3 bytes (01:00:5E)
	multicastMAC = append(multicastMAC, ipAddress[13])         //copy 2. byte of ip into mac -> first byte start at index 12
	multicastMAC = append(multicastMAC, ipAddress[14])         //copy 3. byte of ip into mac -> first byte start at index 12
	multicastMAC = append(multicastMAC, ipAddress[15])         //copy 4. byte of ip into mac -> first byte start at index 12

	//clear msb in 4 octet -> is always 0 in mc mac address
	octet4int := clearBit(int(multicastMAC[3]), 7)
	multicastMAC[3] = byte(octet4int)

	hwAddress := make(net.HardwareAddr, 6)
	copy(hwAddress, multicastMAC)

	return hwAddress
}

// Clears the bit at pos in n.
func clearBit(n int, pos uint) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

func MulticastIPfromMAC(multicastMAC []byte) []string {
	var binaryStringSlice = []string{}

	//write mac bytes as binary strings into slice
	for _, v := range multicastMAC {
		binaryStringSlice = append(binaryStringSlice, fmt.Sprintf("%08b", v))
	}

	binaryString := strings.Join(binaryStringSlice, "")
	last25Bits := binaryString[25:48]
	var missing5Bits = [32]byte{}

	// calculate variable bytes
	for i := 0; i < 32; i++ {
		missing5Bits[i] = byte(i)
	}

	var allMulticastIPStrings = []string{}
	// calculate binary strings of mc ip addresses and add them to slice
	for _, v := range missing5Bits {
		str := fmt.Sprintf("%08b\n", v)[3:8]
		allMulticastIPStrings = append(allMulticastIPStrings, multicastIpPrefix+str+last25Bits)
	}

	return allMulticastIPStrings
}

// calculates the decimal notation of an ip address using a binary string of length 32
func BinaryIpStringToIntString(stringIp string) string {
	//TODO add error handling (no 32 bits, invalid chars etc..)
	octet1 := stringIp[:8]
	octet2 := stringIp[8:16]
	octet3 := stringIp[16:24]
	octet4 := stringIp[24:]

	octet1int := binaryStringToInt(octet1)
	octet2int := binaryStringToInt(octet2)
	octet3int := binaryStringToInt(octet3)
	octet4int := binaryStringToInt(octet4)

	return strconv.Itoa(octet1int) + "." + strconv.Itoa(octet2int) + "." + strconv.Itoa(octet3int) + "." + strconv.Itoa(octet4int)
}

// calculates the integer value of a binary string (containing only 0 and 1)
func binaryStringToInt(binaryString string) int {
	var value = 0
	var exponent = len(binaryString) - 1
	for _, v := range binaryString {
		if string(v) == "1" {
			value += int(math.Pow(2, float64(exponent)))
		}
		exponent--
	}
	return value
}
