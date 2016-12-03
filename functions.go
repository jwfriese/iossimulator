package iossimulator

import (
	"os/exec"
	"time"
)

func IsDeviceAvailable(osString string, deviceTypeString string) error {
	simCtlWrapper := NewSimCtlWrapper()
	environmentParser := NewEnvironmentParser(simCtlWrapper)
	availability := NewSimulatorAvailability(environmentParser)
	return availability.CheckAvailability(osString, deviceTypeString)
}

func CreateDevice(runtime string, deviceType string) (string, error) {
	simCtlWrapper := NewSimCtlWrapper()
	environmentParser := NewEnvironmentParser(simCtlWrapper)
	creator := NewCreator(environmentParser, simCtlWrapper)
	return creator.CreateDevice(runtime, deviceType)
}

func BootDevice(deviceIdentifier string) error {
	simCtlWrapper := NewSimCtlWrapper()
	director := NewDirector(simCtlWrapper)
	return director.BootDevice(deviceIdentifier)
}

func WaitForDeviceToBeReady(deviceIdentifier string) error {
	simCtlWrapper := NewSimCtlWrapper()
	simulatorReadiness := NewSimulatorReadiness(simCtlWrapper)
	isReady, err := simulatorReadiness.IsSimulatorReady(deviceIdentifier)
	for !isReady && err == nil {
		time.Sleep(100 * time.Millisecond)
		isReady, err = simulatorReadiness.IsSimulatorReady(deviceIdentifier)
	}

	return err
}

func ShutdownDevice(deviceIdentifier string) error {
	simCtlWrapper := NewSimCtlWrapper()
	director := NewDirector(simCtlWrapper)
	return director.ShutdownDevice(deviceIdentifier)
}

func DeleteDevice(deviceIdentifier string) error {
	simCtlWrapper := NewSimCtlWrapper()
	director := NewDirector(simCtlWrapper)
	return director.DeleteDevice(deviceIdentifier)
}

func CloseSimulatorApp() error {
	closeCommand := exec.Command("osascript", "-e", "tell app \"Simulator\" to quit")
	return closeCommand.Run()
}
