package iossimulator

type SimulatorEnvironment struct {
	DeviceTypes        map[string]string
	Runtimes           map[string]string
	RuntimeToDeviceMap map[string][]string
}

func (env *SimulatorEnvironment) AvailableRuntimes() []string {
	runtimes := make([]string, 0, len(env.RuntimeToDeviceMap))
	for runtime := range env.RuntimeToDeviceMap {
		runtimes = append(runtimes, runtime)
	}

	return runtimes
}
