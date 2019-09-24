package main

import (
	"fmt"
	`github.com/streamnative/pulsarctl/pkg`
	"os"
)

func main() {
	rootCmd := pkg.NewPulsarctlCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err) // outputs cobra errors
		os.Exit(-1)
	}
}