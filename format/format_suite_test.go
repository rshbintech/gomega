package format_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/velarii/gomega"

	"testing"
)

func TestFormat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Format Suite")
}
