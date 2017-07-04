package iossimulator

import (
	"bufio"
	"bytes"
	"log"
	"strings"
)

type EnvironmentParser interface {
	ParseEnvironment() *SimulatorEnvironment
}

func NewEnvironmentParser(simCtlWrapper SimCtlWrapper) EnvironmentParser {
	return &environmentParser{
		simCtlWrapper: simCtlWrapper,
	}
}

type environmentParser struct {
	simCtlWrapper SimCtlWrapper
}

func (p *environmentParser) ParseEnvironment() *SimulatorEnvironment {
	environment := &SimulatorEnvironment{
		RuntimeToDeviceMap: parseDevices(p.simCtlWrapper),
		DeviceTypes:        parseDeviceTypes(p.simCtlWrapper),
		Runtimes:           parseRuntimes(p.simCtlWrapper),
	}

	return environment
}

func parseDevices(simCtlWrapper SimCtlWrapper) map[string][]string {
	runtimeToDeviceMap := make(map[string][]string)
	infoBytes, err := simCtlWrapper.List()

	if err != nil {
		log.Fatal(err)
	}

	infoBuffer := bytes.NewBuffer(infoBytes)
	infoScanner := bufio.NewScanner(infoBuffer)
	for infoScanner.Scan() {
		if infoScanner.Text() == "== Devices ==" {
			for infoScanner.Text() != "== Device Pairs ==" {
				if strings.HasPrefix(infoScanner.Text(), "--") {
					runtimeString := strings.Trim(infoScanner.Text(), "-\n \t\r")
					devices := []string{}
					infoScanner.Scan()
					hasMoreDevices := !strings.HasPrefix(infoScanner.Text(), "--")
					for hasMoreDevices {
						substrings := strings.SplitN(infoScanner.Text(), "(", 2)
						deviceString := strings.Trim(substrings[0], " \n\t\r")
						devices = append(devices, deviceString)
						infoScanner.Scan()
						hasMoreDevices = !strings.HasPrefix(infoScanner.Text(), "--") && infoScanner.Text() != "== Device Pairs =="

					}

					runtimeToDeviceMap[runtimeString] = devices
				} else {
					infoScanner.Scan()
				}
			}
		}
	}

	return runtimeToDeviceMap
}

func parseDeviceTypes(simCtlWrapper SimCtlWrapper) map[string]string {
	deviceTypes := make(map[string]string)
	infoBytes, err := simCtlWrapper.List()

	if err != nil {
		log.Fatal(err)
	}

	infoBuffer := bytes.NewBuffer(infoBytes)
	infoScanner := bufio.NewScanner(infoBuffer)
	for infoScanner.Scan() {
		if infoScanner.Text() == "== Device Types ==" {
			infoScanner.Scan()
			for infoScanner.Text() != "== Runtimes ==" {
				splitDeviceTypeFromIdResult := strings.SplitN(infoScanner.Text(), "(", 2)
				deviceType := strings.Trim(splitDeviceTypeFromIdResult[0], " \t\n")
				isolateDeviceTypeIdResult := strings.SplitN(splitDeviceTypeFromIdResult[1], ")", 2)
				deviceTypeId := strings.Trim(isolateDeviceTypeIdResult[0], " \t\n")
				deviceTypes[deviceType] = deviceTypeId
				infoScanner.Scan()
			}
		}
	}

	return deviceTypes
}

func parseRuntimes(simCtlWrapper SimCtlWrapper) map[string]string {
	runtimes := make(map[string]string)
	infoBytes, err := simCtlWrapper.List()

	if err != nil {
		log.Fatal(err)
	}

	infoBuffer := bytes.NewBuffer(infoBytes)
	infoScanner := bufio.NewScanner(infoBuffer)
	for infoScanner.Scan() {
		if infoScanner.Text() == "== Runtimes ==" {
			infoScanner.Scan()
			for infoScanner.Text() != "== Devices ==" {
				runtime, parseRuntimeErr := ParseRuntime(infoScanner.Text())
				if parseRuntimeErr != nil {
					log.Fatal(parseRuntimeErr)
				}
				runtimeId, parseRuntimeIdErr := ParseRuntimeId(infoScanner.Text())
				if parseRuntimeIdErr != nil {
					log.Fatal(parseRuntimeIdErr)
				}
				runtimes[runtime] = runtimeId
				infoScanner.Scan()
			}
		}
	}

	return runtimes
}
