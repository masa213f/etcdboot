package main

import (
	"os"

	"github.com/masa213f/etcdboot/cmd/etcdboot/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
