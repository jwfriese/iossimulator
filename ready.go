package iossimulator

import (
	"errors"
	"fmt"
	"strings"
)

type SimulatorReadiness interface {
	IsSimulatorReady(deviceIdentifier string) (bool, error)
}

func NewSimulatorReadiness(simCtlWrapper SimCtlWrapper) SimulatorReadiness {
	return &simulatorReadiness{
		simCtlWrapper: simCtlWrapper,
	}
}

type simulatorReadiness struct {
	simCtlWrapper SimCtlWrapper
}

func (r *simulatorReadiness) IsSimulatorReady(deviceIdentifier string) (bool, error) {
	output, err := r.simCtlWrapper.PrintSpringboardServiceAvailability(deviceIdentifier)
	if err != nil {
		errString := fmt.Sprintf("An error occurred trying to poll simulator with id='%s'. Make sure that it is available as a target using 'xcrun simctl list'.", deviceIdentifier)
		return false, errors.New(errString)
	}

	components := strings.Fields(output)

	return components[2] == "A", nil
}
