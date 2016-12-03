package iossimulator_test

import (
	"errors"

	"github.com/jwfriese/iossimulator"
	"github.com/jwfriese/iossimulator/iossimulatorfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Readiness of simulator", func() {
	var (
		subject           iossimulator.SimulatorReadiness
		fakeSimCtlWrapper *iossimulatorfakes.FakeSimCtlWrapper
	)

	BeforeEach(func() {
		fakeSimCtlWrapper = new(iossimulatorfakes.FakeSimCtlWrapper)
		subject = iossimulator.NewSimulatorReadiness(fakeSimCtlWrapper)
	})

	Describe("When the `simctl` returns an error", func() {
		BeforeEach(func() {
			fakeSimCtlWrapper.PrintSpringboardServiceAvailabilityReturns("", errors.New("some simctlwrapper error"))
		})

		It("returns false", func() {
			isAvailable, _ := subject.IsSimulatorReady("device-uuid")
			Expect(isAvailable).To(BeFalse())
		})

		It("returns an error", func() {
			_, err := subject.IsSimulatorReady("device-uuid")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("An error occurred trying to poll simulator with id='device-uuid'. Make sure that it is available as a target using 'xcrun simctl list'."))
		})
	})

	Describe("When the `simctl` returns a value indicating that the sim is not ready yet", func() {
		BeforeEach(func() {
			fakeSimCtlWrapper.PrintSpringboardServiceAvailabilityReturns("0x1be07    M   D   com.apple.springboard.services", nil)
		})

		It("returns false", func() {
			isAvailable, _ := subject.IsSimulatorReady("device-uuid")
			Expect(isAvailable).To(BeFalse())
		})

		It("returns no error", func() {
			_, err := subject.IsSimulatorReady("device-uuid")
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("When the `simctl` returns a value indicating that the sim is ready", func() {
		BeforeEach(func() {
			fakeSimCtlWrapper.PrintSpringboardServiceAvailabilityReturns("0x1be07    M   A   com.apple.springboard.services", nil)
		})

		It("returns true", func() {
			isAvailable, _ := subject.IsSimulatorReady("device-uuid")
			Expect(isAvailable).To(BeTrue())
		})

		It("returns no error", func() {
			_, err := subject.IsSimulatorReady("device-uuid")
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
