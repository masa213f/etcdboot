package cmd

import (
	"github.com/masa213f/etcdboot"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:          "etcdboot",
	Short:        "A deploy tool of etcd cluster",
	Version:      etcdboot.Version,
	SilenceUsage: true,
}

func init() {
	RootCmd.AddCommand(serviceCreateCmd)
	RootCmd.AddCommand(serviceDeleteCmd)
	RootCmd.AddCommand(serviceUpdateCmd)
}
