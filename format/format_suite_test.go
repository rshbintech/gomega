package format_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/rshbintech/gomega"

	"testing"
)

func TestFormat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Format Suite")
}
