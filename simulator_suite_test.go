package iossimulator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSimulator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "iOS Simulator Suite")
}
