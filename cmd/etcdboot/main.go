package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/masa213f/etcdboot/pkg/config"
	"github.com/masa213f/etcdboot/pkg/systemd"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("./etcdboot SERVICE_NAME MEMBER_NAME PATH_TO_CLUSTER_YAML")
		os.Exit(1)
	}

	cluster, err := config.ReadClusterConfig(args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = systemd.CreateServiceFile(args[0], args[1], cluster)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("ok")
}
