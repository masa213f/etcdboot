package cmd

import (
	"errors"
	"fmt"

	"github.com/masa213f/etcdboot/pkg/config"
	"github.com/masa213f/etcdboot/pkg/generator"
	"github.com/spf13/cobra"
)

var serviceCreateCmd = &cobra.Command{
	Use:   "create [SERVICE_NAME]",
	Short: "Create a systemd service",
	Args:  cobra.MaximumNArgs(1),
	RunE:  serviceCreateMain,
}

var serviceCreateCmdOption struct {
	clusterFile string
	etcdBinary  string
	memberName  string
}

func init() {
	serviceCreateCmd.Flags().StringVarP(&serviceCreateCmdOption.clusterFile, "cluster-file", "c", "", "path to cluster.yaml")
	serviceCreateCmd.Flags().StringVarP(&serviceCreateCmdOption.etcdBinary, "etcd-binary", "b", "/usr/local/bin/etcd", "path to etcd binary")
	serviceCreateCmd.Flags().StringVarP(&serviceCreateCmdOption.memberName, "member-name", "m", "", "etcd member name")
}

func serviceCreateMain(cmd *cobra.Command, args []string) error {
	serviceName := defaultServiceName
	if len(args) > 0 {
		serviceName = args[0]
	}

	var cluster *config.EtcdCluster
	if serviceCreateCmdOption.clusterFile == "" {
		// Fix me. Let's use default setting.
		return errors.New("cluster-file is not specified")
	}
	cluster, err := config.ReadClusterConfig(serviceCreateCmdOption.clusterFile)
	if err != nil {
		return err
	}
	if serviceCreateCmdOption.memberName == "" {
		if len(cluster.Members) != 1 {
			return errors.New("member-name is not specified")
		}
		serviceCreateCmdOption.memberName = cluster.Members[0].Name
	}

	serviceFilePath := getServiceFilePath(serviceName)
	err = createServiceFile(serviceFilePath, &generator.ServiceFileOption{
		Cluster:     cluster,
		MemberName:  serviceCreateCmdOption.memberName,
		EtcdBinPath: serviceCreateCmdOption.etcdBinary,
		DataDirPath: getDataDirPath(serviceName),
	})
	if err != nil {
		return err
	}
	fmt.Printf("create service file: %s\n", serviceFilePath)

	return systemdReloadAndStart(serviceName)
}
