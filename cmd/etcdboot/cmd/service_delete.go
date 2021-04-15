package cmd

import (
	"fmt"
	"os"

	"github.com/masa213f/etcdboot/pkg/systemd"
	"github.com/spf13/cobra"
)

var serviceDeleteCmd = &cobra.Command{
	Use:   "delete [SERVICE_NAME]",
	Short: "Delete a systemd service",
	Args:  cobra.MaximumNArgs(1),
	RunE:  serviceDeleteMain,
}

var serviceDeleteCmdOption struct {
	daleteData bool
}

func init() {
	serviceDeleteCmd.Flags().BoolVar(&serviceDeleteCmdOption.daleteData, "data", false, "")
}

func serviceDeleteMain(cmd *cobra.Command, args []string) error {
	serviceName := defaultServiceName
	if len(args) > 0 {
		serviceName = args[0]
	}

	serviceFilePath := getServiceFilePath(serviceName)
	if !fileExists(serviceFilePath) {
		return fmt.Errorf("file does not exist: %s", serviceFilePath)
	}
	err := systemd.ServiceStop(serviceName)
	if err != nil {
		return fmt.Errorf("failed to stop service: err=%v", err)
	}
	err = systemd.ServiceDisable(serviceName)
	if err != nil {
		return fmt.Errorf("failed to disable service: err=%v", err)
	}

	if serviceDeleteCmdOption.daleteData {
		dataDir := getDataDirPath(serviceName)
		err := os.RemoveAll(dataDir)
		if err != nil {
			return fmt.Errorf("failed to delete data dir: err=%v", err)
		}
	}

	err = os.Remove(serviceFilePath)
	if err != nil {
		return fmt.Errorf("failed to delete service file: err=%v", err)
	}

	err = systemd.DaemonReload()
	if err != nil {
		return fmt.Errorf("failed to reload daemon: err=%v", err)
	}
	return nil
}
