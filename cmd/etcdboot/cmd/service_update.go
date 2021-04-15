package cmd

import (
	"errors"
	"fmt"

	"github.com/masa213f/etcdboot/pkg/config"
	"github.com/masa213f/etcdboot/pkg/generator"
	"github.com/spf13/cobra"
)

var serviceUpdateCmd = &cobra.Command{
	Use:   "update [SERVICE_NAME]",
	Short: "Update a systemd service",
	Args:  cobra.MaximumNArgs(1),
	RunE:  serviceUpdateMain,
}

var serviceUpdateCmdOption struct {
	clusterFile string
	etcdBinary  string
	memberName  string
}

func init() {
	serviceUpdateCmd.Flags().StringVarP(&serviceUpdateCmdOption.clusterFile, "cluster-file", "c", "", "path to cluster.yaml")
	serviceUpdateCmd.Flags().StringVarP(&serviceUpdateCmdOption.etcdBinary, "etcd-binary", "b", "/usr/local/bin/etcd", "path to etcd binary")
	serviceUpdateCmd.Flags().StringVarP(&serviceUpdateCmdOption.memberName, "member-name", "m", "", "etcd member name")
}

func serviceUpdateMain(cmd *cobra.Command, args []string) error {
	serviceName := defaultServiceName
	if len(args) > 0 {
		serviceName = args[0]
	}

	var cluster *config.EtcdCluster
	if serviceUpdateCmdOption.clusterFile == "" {
		// Fix me. Let's use default setting.
		return errors.New("cluster-file is not specified")
	}
	cluster, err := config.ReadClusterConfig(serviceUpdateCmdOption.clusterFile)
	if err != nil {
		return err
	}
	if serviceUpdateCmdOption.memberName == "" {
		if len(cluster.Members) != 1 {
			return errors.New("member-name is not specified")
		}
		serviceUpdateCmdOption.memberName = cluster.Members[0].Name
	}

	serviceFilePath := getServiceFilePath(serviceName)
	err = updateServiceFile(serviceFilePath, &generator.ServiceFileOption{
		Cluster:     cluster,
		MemberName:  serviceUpdateCmdOption.memberName,
		EtcdBinPath: serviceUpdateCmdOption.etcdBinary,
		DataDirPath: getDataDirPath(serviceName),
	})
	if err != nil {
		return err
	}
	fmt.Printf("create service file: %s\n", serviceFilePath)

	return systemdReloadAndStart(serviceName)
}
