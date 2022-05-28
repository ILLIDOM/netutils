package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MulticastIPfromMAC(multicastMAC []byte) {
	//TODO -> implement generation of multicast ips from mac
	fmt.Println("calculation multicast IPs")
	var binaryStringSlice = []string{}

	var first4bits = "1110"

	for _, v := range multicastMAC {
		binaryStringSlice = append(binaryStringSlice, fmt.Sprintf("%08b", v))
	}

	binaryString := strings.Join(binaryStringSlice, "")
	fmt.Println(binaryString)
	last25Bits := binaryString[25:48]
	fmt.Println(last25Bits)

	var missing5Bits = [32]byte{}

	for i := 0; i < 32; i++ {
		missing5Bits[i] = byte(i)
	}

	var allMulticastIPStrings = [32]string{}

	for i, v := range missing5Bits {
		str := fmt.Sprintf("%08b\n", v)[3:8]
		allMulticastIPStrings[i] = first4bits + str + last25Bits
	}

	fmt.Println(allMulticastIPStrings)
	fmt.Println(binaryIpStringToIntString(allMulticastIPStrings[0]))
}

func binaryIpStringToIntString(stringIp string) string {
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
