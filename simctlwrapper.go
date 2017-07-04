package iossimulator

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type SimCtlWrapper interface {
	List() ([]byte, error)
	Create(name string, runtime string, deviceType string) (string, error)
	Boot(deviceIdentifier string) error
	Shutdown(deviceIdentifier string) error
	Delete(deviceIdentifier string) error
	PrintSpringboardServiceAvailability(deviceIdentifier string) (string, error)
}

func NewSimCtlWrapper() SimCtlWrapper {
	return &simCtlWrapper{}
}

type simCtlWrapper struct{}

func (w *simCtlWrapper) List() ([]byte, error) {
	getInfoCommand := exec.Command("xcrun", "simctl", "list")
	combinedOutput, err := getInfoCommand.CombinedOutput()
	if err != nil {
		return []byte(`Could not list`), errors.New(fmt.Sprintf("Error: %v\nFull output: %v", err, combinedOutput))
	}

	return combinedOutput, nil
}

func (w *simCtlWrapper) Create(name string, deviceType string, runtime string) (string, error) {
	createCommand := exec.Command("xcrun", "simctl", "create", name, deviceType, runtime)
	combinedOutput, err := createCommand.CombinedOutput()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error: %v\nFull output: %v", err, string(combinedOutput)))
	}
	untrimmedDeviceUuid := string(combinedOutput)
	deviceUuid := strings.Trim(untrimmedDeviceUuid, " \t\n")
	return deviceUuid, nil
}

func (w *simCtlWrapper) Boot(deviceIdentifier string) error {
	bootCommand := exec.Command("open", "-a", "Simulator", "--args", "-CurrentDeviceUDID", deviceIdentifier)
	combinedOutput, err := bootCommand.CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprintf("Error: %v\nFull output: %v", err, string(combinedOutput)))
	}

	return nil
}

func (w *simCtlWrapper) Shutdown(deviceIdentifier string) error {
	shutdownCommand := exec.Command("xcrun", "simctl", "shutdown", deviceIdentifier)
	combinedOutput, err := shutdownCommand.CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprintf("Error: %v\nFull output: %v", err, string(combinedOutput)))
	}

	return nil
}

func (w *simCtlWrapper) Delete(deviceIdentifier string) error {
	deleteCommand := exec.Command("xcrun", "simctl", "delete", deviceIdentifier)
	combinedOutput, err := deleteCommand.CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprintf("Error: %v\nFull output: %v", err, string(combinedOutput)))
	}

	return nil
}

func (w *simCtlWrapper) PrintSpringboardServiceAvailability(deviceIdentifier string) (string, error) {
	command := exec.Command("xcrun", "simctl", "spawn", deviceIdentifier, "launchctl", "print", "system", "|", "grep", "com.apple.springboard.services")
	bytes, err := command.Output()
	untrimmedResult := string(bytes)
	result := strings.Trim(untrimmedResult, " \t\n")
	return result, err
}
