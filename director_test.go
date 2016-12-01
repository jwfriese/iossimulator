package iossimulator_test

import (
	"errors"

	"github.com/jwfriese/iossimulator"
	"github.com/jwfriese/iossimulator/iossimulatorfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Director", func() {
	var (
		fakeSimCtlWrapper *iossimulatorfakes.FakeSimCtlWrapper
		subject           iossimulator.Director
	)

	BeforeEach(func() {
		fakeSimCtlWrapper = new(iossimulatorfakes.FakeSimCtlWrapper)
		subject = iossimulator.NewDirector(fakeSimCtlWrapper)
	})

	Describe("Booting a simulator device", func() {
		Context("When the attempt to boot fails", func() {
			It("returns an error that includes the error from 'simctl'", func() {
				errString := "simctl boot error"
				fakeSimCtlWrapper.BootReturns(errors.New(errString))

				err := subject.BootDevice("device-id")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Error booting device: simctl boot error"))
			})
		})

		Context("When the attempt to boot succeeds", func() {
			It("returns no error", func() {
				fakeSimCtlWrapper.BootReturns(nil)

				err := subject.BootDevice("device-id")
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Describe("Shutting down a simulator device", func() {
		Context("When the attempt to shutdown fails", func() {
			It("returns an error that includes the error from 'simctl'", func() {
				errString := "simctl shutdown error"
				fakeSimCtlWrapper.ShutdownReturns(errors.New(errString))

				err := subject.ShutdownDevice("device-id")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Error shutting down device: simctl shutdown error"))
			})
		})

		Context("When the attempt to shutdown succeeds", func() {
			It("returns no error", func() {
				fakeSimCtlWrapper.ShutdownReturns(nil)

				err := subject.ShutdownDevice("device-id")
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Describe("Deleting a simulator device", func() {
		Context("When the attempt to delete fails", func() {
			It("returns an error that includes the error from 'simctl'", func() {
				errString := "simctl delete error"
				fakeSimCtlWrapper.DeleteReturns(errors.New(errString))

				err := subject.DeleteDevice("device-id")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Error deleting device: simctl delete error"))
			})
		})

		Context("When the attempt to delete succeeds", func() {
			It("returns no error", func() {
				fakeSimCtlWrapper.DeleteReturns(nil)

				err := subject.DeleteDevice("device-id")
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
