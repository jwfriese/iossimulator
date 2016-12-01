package stringslice_test

import (
	"github.com/jwfriese/iossimulator/stringslice"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Contains", func() {
	It("returns false when the slice does not contain the string", func() {
		slice := []string{"turtle", "puppy", "kitten"}
		Expect(stringslice.Contains(slice, "crab")).To(BeFalse())
	})

	It("returns true when the slice does not contain the string", func() {
		slice := []string{"turtle", "puppy", "kitten"}
		Expect(stringslice.Contains(slice, "puppy")).To(BeTrue())
	})
})
