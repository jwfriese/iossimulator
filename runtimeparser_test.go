package iossimulator_test

import (
	"github.com/jwfriese/iossimulator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parsing runtime data from the output of `xcrun simctl list`", func() {
	Describe("<version> (meta) (id) format", func() {
		var iOSInput string
		var watchOSInput string
		var tvOSInput string

		BeforeEach(func() {
			iOSInput = "iOS 9.3 (9.3 - 13E233) (com.apple.CoreSimulator.SimRuntime.iOS-9-3)"
			watchOSInput = "watchOS 4.0 (4.0 - 15R5307e) (com.apple.CoreSimulator.SimRuntime.watchOS-4-0)"
			tvOSInput = "tvOS 11.0 (11.0 - 15J5310e) (com.apple.CoreSimulator.SimRuntime.tvOS-11-0)"
		})

		It("can parse the runtime", func() {
			result, err := iossimulator.ParseRuntime(iOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("iOS 9.3"))

			result, err = iossimulator.ParseRuntime(watchOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("watchOS 4.0"))

			result, err = iossimulator.ParseRuntime(tvOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("tvOS 11.0"))
		})

		It("can parse the runtime id", func() {
			result, err := iossimulator.ParseRuntimeId(iOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("com.apple.CoreSimulator.SimRuntime.iOS-9-3"))

			result, err = iossimulator.ParseRuntimeId(watchOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("com.apple.CoreSimulator.SimRuntime.watchOS-4-0"))

			result, err = iossimulator.ParseRuntimeId(tvOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("com.apple.CoreSimulator.SimRuntime.tvOS-11-0"))
		})
	})

	Describe("<version> (meta) - id format", func() {
		var iOSInput string
		var watchOSInput string
		var tvOSInput string

		BeforeEach(func() {
			iOSInput = "iOS 9.3 (9.3 - 13E233) - com.apple.CoreSimulator.SimRuntime.iOS-9-3"
			watchOSInput = "watchOS 4.0 (4.0 - 15R5307e) - com.apple.CoreSimulator.SimRuntime.watchOS-4-0"
			tvOSInput = "tvOS 11.0 (11.0 - 15J5310e) - com.apple.CoreSimulator.SimRuntime.tvOS-11-0"
		})

		It("can parse the runtime", func() {
			result, err := iossimulator.ParseRuntime(iOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("iOS 9.3"))

			result, err = iossimulator.ParseRuntime(watchOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("watchOS 4.0"))

			result, err = iossimulator.ParseRuntime(tvOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("tvOS 11.0"))
		})

		It("can parse the runtime id", func() {
			result, err := iossimulator.ParseRuntimeId(iOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("com.apple.CoreSimulator.SimRuntime.iOS-9-3"))

			result, err = iossimulator.ParseRuntimeId(watchOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("com.apple.CoreSimulator.SimRuntime.watchOS-4-0"))

			result, err = iossimulator.ParseRuntimeId(tvOSInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal("com.apple.CoreSimulator.SimRuntime.tvOS-11-0"))
		})
	})
})
