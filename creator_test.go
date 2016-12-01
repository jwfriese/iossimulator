package iossimulator_test

import (
	"errors"

	"github.com/jwfriese/iossimulator"
	"github.com/jwfriese/iossimulator/iossimulatorfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creator", func() {
	var (
		fakeEnvironmentParser *iossimulatorfakes.FakeEnvironmentParser
		fakeSimCtlWrapper     *iossimulatorfakes.FakeSimCtlWrapper
		subject               iossimulator.Creator
	)

	BeforeEach(func() {
		fakeEnvironmentParser = new(iossimulatorfakes.FakeEnvironmentParser)
		fakeSimCtlWrapper = new(iossimulatorfakes.FakeSimCtlWrapper)
		subject = iossimulator.NewCreator(fakeEnvironmentParser, fakeSimCtlWrapper)
	})

	Describe("CreateDevice", func() {
		var (
			newDeviceUuid string
			err           error
		)

		BeforeEach(func() {
			environment := &iossimulator.SimulatorEnvironment{
				DeviceTypes: map[string]string{
					"iPhone 4s": "iPhone-4s-identifier",
					"iPhone 6":  "iPhone-6-identifier",
				},
				Runtimes: map[string]string{
					"iOS 9.0":  "iOS-9-0-identifier",
					"iOS 10.0": "iOS-10-0-identifier",
				},
			}

			fakeEnvironmentParser.ParseEnvironmentReturns(environment)
		})

		Context("When the requested runtime is unavailable", func() {
			BeforeEach(func() {
				newDeviceUuid, err = subject.CreateDevice("iOS 10.1", "iPhone 6")
			})

			It("returns no uuid", func() {
				Expect(newDeviceUuid).To(Equal(""))
			})

			It("returns an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Could not create simulator device: 'iOS 10.1' runtime is not available"))
			})
		})

		Context("When the requested device type is unavailable", func() {
			BeforeEach(func() {
				newDeviceUuid, err = subject.CreateDevice("iOS 10.0", "iPhone 7")
			})

			It("returns no uuid", func() {
				Expect(newDeviceUuid).To(Equal(""))
			})

			It("returns an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Could not create simulator device: 'iPhone 7' device type is not available"))
			})
		})

		Context("When both the runtime and device type are available", func() {
			Context("When the SimCtlWrapper errors trying to create the device", func() {
				BeforeEach(func() {
					errString := "test simctlwrapper error"
					fakeSimCtlWrapper.CreateReturns("", errors.New(errString))
					newDeviceUuid, err = subject.CreateDevice("iOS 10.0", "iPhone 6")
				})

				It("returns no uuid", func() {
					Expect(newDeviceUuid).To(Equal(""))
				})

				It("returns an error", func() {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("Could not create simulator device: test simctlwrapper error"))
				})
			})

			Context("When the SimCtlWrapper successfully creates a new dimulator device", func() {
				BeforeEach(func() {
					fakeSimCtlWrapper.CreateReturns("new-device-uuid", nil)
					newDeviceUuid, err = subject.CreateDevice("iOS 10.0", "iPhone 6")
				})

				It("returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("returns the UUID of the newly-created simulator device", func() {
					Expect(newDeviceUuid).To(Equal("new-device-uuid"))
				})
			})
		})
	})
})
