package gbytes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/rshbintech/gomega"

	"testing"
)

func TestGbytes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gbytes Suite")
}
