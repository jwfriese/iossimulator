package iossimulator

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
