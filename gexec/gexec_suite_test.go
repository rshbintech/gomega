package gexec_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/velarii/gomega"
	"github.com/velarii/gomega/gexec"

	"testing"
)

func TestGexec(t *testing.T) {
	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Gexec Suite")
}
