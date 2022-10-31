package matchers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/velarii/gomega"
	. "github.com/velarii/gomega/matchers"
)

var _ = Describe("BeFalse", func() {
	It("should handle true and false correctly", func() {
		Expect(true).ShouldNot(BeFalse())
		Expect(false).Should(BeFalse())
	})

	It("should only support booleans", func() {
		success, err := (&BeFalseMatcher{}).Match("foo")
		Expect(success).Should(BeFalse())
		Expect(err).Should(HaveOccurred())
	})
})
