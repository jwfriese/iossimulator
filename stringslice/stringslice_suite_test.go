package stringslice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStringslice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringSlice Suite")
}
