package main

import (
	"github.com/ILLIDOM/neturils/multicast/cmd"

	"fmt"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
