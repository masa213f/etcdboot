package systemd

import (
	"os/exec"
)

func DaemonReload() error {
	_, err := exec.Command("systemctl", "-q", "daemon-reload").Output()
	return err
}

func ServiceEnable(serviceName string) error {
	_, err := exec.Command("systemctl", "-q", "enable", serviceName).Output()
	return err
}

func ServiceDisable(serviceName string) error {
	_, err := exec.Command("systemctl", "-q", "disable", serviceName).Output()
	return err
}

func ServiceStart(serviceName string) error {
	_, err := exec.Command("systemctl", "-q", "start", serviceName).Output()
	return err
}

func ServiceStop(serviceName string) error {
	_, err := exec.Command("systemctl", "-q", "stop", serviceName).Output()
	return err
}
