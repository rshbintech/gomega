package bipartitegraph_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/rshbintech/gomega"

	"testing"
)

func TestBipartitegraph(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bipartitegraph Suite")
}
