package iossimulator

import (
	"errors"
	"fmt"
)

type Creator interface {
	CreateDevice(runtime string, deviceType string) (string, error)
}

func NewCreator(environmentParser EnvironmentParser, simCtlWrapper SimCtlWrapper) Creator {
	return &creator{
		environmentParser: environmentParser,
		simCtlWrapper:     simCtlWrapper,
	}
}

type creator struct {
	environmentParser EnvironmentParser
	simCtlWrapper     SimCtlWrapper
}

func (c *creator) CreateDevice(runtime string, deviceType string) (string, error) {
	environment := c.environmentParser.ParseEnvironment()
	var runtimeId string
	var ok bool
	if runtimeId, ok = environment.Runtimes[runtime]; !ok {
		errString := fmt.Sprintf("Could not create simulator device: '%s' runtime is not available", runtime)
		return "", errors.New(errString)
	}

	var deviceTypeId string
	if deviceTypeId, ok = environment.DeviceTypes[deviceType]; !ok {
		errString := fmt.Sprintf("Could not create simulator device: '%s' device type is not available", deviceType)
		return "", errors.New(errString)
	}

	newDeviceUuid, err := c.simCtlWrapper.Create("iossimulator-device", deviceTypeId, runtimeId)
	if err != nil {
		errString := fmt.Sprintf("Could not create simulator device: %s", err.Error())
		return "", errors.New(errString)
	}

	return newDeviceUuid, nil
}
