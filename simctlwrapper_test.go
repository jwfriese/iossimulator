package iossimulator_test

import (
	"time"

	"github.com/jwfriese/iossimulator"
	"github.com/jwfriese/iossimulator/stringslice"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SimCtlWrapper", func() {
	var subject iossimulator.SimCtlWrapper
	var parser iossimulator.EnvironmentParser

	BeforeEach(func() {
		subject = iossimulator.NewSimCtlWrapper()
		parser = iossimulator.NewEnvironmentParser(subject)
	})

	It("can create, boot, shutdown, and delete a device", func() {
		simulatorEnvironment := parser.ParseEnvironment()
		devices := simulatorEnvironment.RuntimeToDeviceMap["iOS 10.0"]
		Expect(stringslice.Contains(devices, "iossimulator")).To(BeFalse())

		newDeviceUuid, err := subject.Create("iossimulator", "iPhone 6", "com.apple.CoreSimulator.SimRuntime.iOS-10-0")
		Expect(err).ToNot(HaveOccurred())
		Expect(newDeviceUuid).ToNot(Equal(""))

		simulatorEnvironment = parser.ParseEnvironment()
		devices = simulatorEnvironment.RuntimeToDeviceMap["iOS 10.0"]
		Expect(stringslice.Contains(devices, "iossimulator")).To(BeTrue())

		err = subject.Boot(newDeviceUuid)
		Expect(err).ToNot(HaveOccurred())

		time.Sleep(20000 * time.Millisecond)

		err = subject.Shutdown(newDeviceUuid)
		Expect(err).ToNot(HaveOccurred())

		err = subject.Delete(newDeviceUuid)
		Expect(err).ToNot(HaveOccurred())

		simulatorEnvironment = parser.ParseEnvironment()
		devices = simulatorEnvironment.RuntimeToDeviceMap["iOS 10.0"]
		Expect(stringslice.Contains(devices, "iossimulator")).To(BeFalse())
	})
})
