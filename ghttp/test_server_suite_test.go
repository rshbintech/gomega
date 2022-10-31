package ghttp_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/velarii/gomega"

	"testing"
)

func TestGHTTP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GHTTP Suite")
}
