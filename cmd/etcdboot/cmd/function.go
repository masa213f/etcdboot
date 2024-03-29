package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/masa213f/etcdboot/pkg/generator"
	"github.com/masa213f/etcdboot/pkg/systemd"
)

const defaultServiceName = "etcd"

func getServiceFilePath(serviceName string) string {
	return filepath.Join("/etc/systemd/system/", serviceName+".service")
}

func getDataDirPath(serviceName string) string {
	return filepath.Join("/var/lib", serviceName)
}

// checkOverwritableService confirms the service is overwritable or not.
// All service files created by etcdboot has the specific first line.
// So this function check the first line.
func checkOverwritableService(serviceFilePath string) error {
	f, err := os.Open(serviceFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	firstLine := scanner.Text()
	if firstLine != "# Generated by etcdboot; DO NOT EDIT." {
		return errors.New("not overwritable")
	}
	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func createServiceFile(serviceFilePath string, opt *generator.ServiceFileOption) error {
	if fileExists(serviceFilePath) {
		return fmt.Errorf("file exists: %s", serviceFilePath)
	}

	content, err := generator.ServiceFileContent(opt)
	if err != nil {
		return err
	}
	return os.WriteFile(serviceFilePath, content, 0644)
}

func updateServiceFile(serviceFilePath string, opt *generator.ServiceFileOption) error {
	if !fileExists(serviceFilePath) {
		return fmt.Errorf("file does not exist: %s", serviceFilePath)
	}

	content, err := generator.ServiceFileContent(opt)
	if err != nil {
		return err
	}
	return os.WriteFile(serviceFilePath, content, 0644)
}

func systemdReloadAndStart(serviceName string) error {
	err := systemd.DaemonReload()
	if err != nil {
		return fmt.Errorf("failed to reload daemon: err=%v", err)
	}
	err = systemd.ServiceEnable(serviceName)
	if err != nil {
		return fmt.Errorf("failed to enable service: err=%v", err)
	}
	err = systemd.ServiceStart(serviceName)
	if err != nil {
		return fmt.Errorf("failed to start service: err=%v", err)
	}
	return nil
}
