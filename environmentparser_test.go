package iossimulator_test

import (
	"github.com/jwfriese/iossimulator"
	"github.com/jwfriese/iossimulator/iossimulatorfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EnvironmentParser", func() {
	var (
		fakeSimCtlWrapper *iossimulatorfakes.FakeSimCtlWrapper
		subject           iossimulator.EnvironmentParser
	)

	BeforeEach(func() {
		fakeSimCtlWrapper = new(iossimulatorfakes.FakeSimCtlWrapper)
		fakeSimCtlWrapper.ListReturns(
			[]byte("== Device Types ==\n"+
				"iPhone 5 (iPhone-5-device-type-id)\n"+
				"iPhone 6s (iPhone-6s-device-type-id)\n"+
				"== Runtimes ==\n"+
				"iOS 9.3 () (com.apple.iOS-9-3-runtime-id)\n"+
				"iOS 10.0 () (com.apple.iOS-10-0-runtime-id)\n"+
				"== Devices ==\n"+
				"-- iOS 9.1 --\n"+
				"	iPhone 4s (iPhone-4s-sim-id) (iPhone4s-sim-state)\n"+
				"	iPhone 6 (iPhone-6-sim-id) (iPhone-6-sim-state)\n"+
				"-- iOS 10.0 --\n"+
				"	iPhone 5 (iPhone-5-sim-id) (iPhone5-sim-state)\n"+
				"== Device Pairs =="),
			nil)
		subject = iossimulator.NewEnvironmentParser(fakeSimCtlWrapper)
	})

	It("parses devices from `xcrun simctl list` into the sim environment model", func() {
		environment := subject.ParseEnvironment()
		Expect(environment).ToNot(BeNil())

		nineDotOneDevices := environment.RuntimeToDeviceMap["iOS 9.1"]
		Expect(nineDotOneDevices).To(Equal([]string{"iPhone 4s", "iPhone 6"}))

		tenDevices := environment.RuntimeToDeviceMap["iOS 10.0"]
		Expect(tenDevices).To(Equal([]string{"iPhone 5"}))

		notThereDevices := environment.RuntimeToDeviceMap["iOS 8.0"]
		Expect(notThereDevices).To(BeNil())
	})

	It("parses device types from `xcrun simctl list` into the sim environment model", func() {
		environment := subject.ParseEnvironment()
		Expect(environment).ToNot(BeNil())

		fiveDeviceTypeId := environment.DeviceTypes["iPhone 5"]
		Expect(fiveDeviceTypeId).To(Equal("iPhone-5-device-type-id"))

		sixEssDeviceTypeId := environment.DeviceTypes["iPhone 6s"]
		Expect(sixEssDeviceTypeId).To(Equal("iPhone-6s-device-type-id"))

		notThereDeviceType := environment.DeviceTypes["iPhone 17"]
		Expect(notThereDeviceType).To(Equal(""))
	})

	It("parses runtimes from `xcrun simctl list` into the sim environment model", func() {
		environment := subject.ParseEnvironment()
		Expect(environment).ToNot(BeNil())

		nineThreeRuntimeId := environment.Runtimes["iOS 9.3"]
		Expect(nineThreeRuntimeId).To(Equal("com.apple.iOS-9-3-runtime-id"))

		tenRuntimeId := environment.Runtimes["iOS 10.0"]
		Expect(tenRuntimeId).To(Equal("com.apple.iOS-10-0-runtime-id"))

		notThereRuntime := environment.Runtimes["iOS 15.8"]
		Expect(notThereRuntime).To(Equal(""))
	})
})
