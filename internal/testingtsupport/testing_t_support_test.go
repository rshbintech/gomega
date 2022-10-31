package testingtsupport_test

import (
	. "github.com/velarii/gomega"

	"testing"
)

func TestTestingT(t *testing.T) {
	RegisterTestingT(t)
	Ω(true).Should(BeTrue())
}

func TestNewGomegaWithT(t *testing.T) {
	g := NewWithT(t)
	g.Expect(true).To(BeTrue())
}
