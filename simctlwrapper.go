package iossimulator

import (
	"os/exec"
	"strings"
)

type SimCtlWrapper interface {
	List() ([]byte, error)
	Create(name string, runtime string, deviceType string) (string, error)
	Boot(deviceIdentifier string) error
	Shutdown(deviceIdentifier string) error
	Delete(deviceIdentifier string) error
}

func NewSimCtlWrapper() SimCtlWrapper {
	return &simCtlWrapper{}
}

type simCtlWrapper struct{}

func (w *simCtlWrapper) List() ([]byte, error) {
	getInfoCommand := exec.Command("xcrun", "simctl", "list")
	return getInfoCommand.Output()
}

func (w *simCtlWrapper) Create(name string, deviceType string, runtime string) (string, error) {
	createCommand := exec.Command("xcrun", "simctl", "create", name, deviceType, runtime)
	bytes, err := createCommand.Output()
	untrimmedDeviceUuid := string(bytes)
	deviceUuid := strings.Trim(untrimmedDeviceUuid, " \t\n")
	return deviceUuid, err
}

func (w *simCtlWrapper) Boot(deviceIdentifier string) error {
	bootCommand := exec.Command("xcrun", "simctl", "boot", deviceIdentifier)
	_, err := bootCommand.Output()
	return err
}

func (w *simCtlWrapper) Shutdown(deviceIdentifier string) error {
	shutdownCommand := exec.Command("xcrun", "simctl", "shutdown", deviceIdentifier)
	_, err := shutdownCommand.Output()
	return err
}

func (w *simCtlWrapper) Delete(deviceIdentifier string) error {
	deleteCommand := exec.Command("xcrun", "simctl", "delete", deviceIdentifier)
	_, err := deleteCommand.Output()
	return err
}
