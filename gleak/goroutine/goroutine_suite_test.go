package goroutine

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/rshbintech/gomega"
)

func TestGoroutine(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goroutine Suite")
}
