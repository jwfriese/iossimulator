package iossimulator

import (
	"errors"
	"fmt"
)

type Director interface {
	BootDevice(simulatorUuid string) error
	ShutdownDevice(simulatorUuid string) error
	DeleteDevice(simulatorUuid string) error
}

func NewDirector(simCtlWrapper SimCtlWrapper) Director {
	return &director{
		simCtlWrapper: simCtlWrapper,
	}
}

type director struct {
	simCtlWrapper SimCtlWrapper
}

func (d *director) BootDevice(simulatorUUID string) error {
	err := d.simCtlWrapper.Boot(simulatorUUID)
	if err != nil {
		errString := fmt.Sprintf("Error booting device: %s", err.Error())
		return errors.New(errString)
	}

	return nil
}

func (d *director) ShutdownDevice(simulatorUUID string) error {
	err := d.simCtlWrapper.Shutdown(simulatorUUID)
	if err != nil {
		errString := fmt.Sprintf("Error shutting down device: %s", err.Error())
		return errors.New(errString)
	}

	return nil
}

func (d *director) DeleteDevice(simulatorUUID string) error {
	err := d.simCtlWrapper.Delete(simulatorUUID)
	if err != nil {
		errString := fmt.Sprintf("Error deleting device: %s", err.Error())
		return errors.New(errString)
	}

	return nil
}
