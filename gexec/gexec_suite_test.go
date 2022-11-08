package gexec_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/rshbintech/gomega"
	"github.com/rshbintech/gomega/gexec"

	"testing"
)

func TestGexec(t *testing.T) {
	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Gexec Suite")
}
