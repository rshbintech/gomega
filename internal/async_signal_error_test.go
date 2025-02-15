package internal_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/rshbintech/gomega"
	"github.com/rshbintech/gomega/internal"
)

var _ = Describe("AsyncSignalError", func() {
	Describe("StopTrying", func() {
		Describe("building StopTrying errors", func() {
			It("returns a correctly configured StopTrying error", func() {
				st := StopTrying("I've tried 17 times - give up!")
				Ω(st.Error()).Should(Equal("I've tried 17 times - give up!"))
				Ω(errors.Unwrap(st)).Should(BeNil())
				Ω(st.(*internal.AsyncSignalErrorImpl).IsStopTrying()).Should(BeTrue())
			})
		})

		Describe("Wrapping other errors", func() {
			It("can wrap other errors", func() {
				st := StopTrying("Welp! Time to give up")
				Ω(st.Error()).Should(Equal("Welp! Time to give up"))
				st = st.Wrap(fmt.Errorf("ERR_GIVE_UP"))
				Ω(errors.Unwrap(st)).Should(Equal(fmt.Errorf("ERR_GIVE_UP")))
				Ω(st.Error()).Should(Equal("Welp! Time to give up: ERR_GIVE_UP"))
			})
		})

		Describe("When attaching objects", func() {
			It("attaches them, with their descriptions", func() {
				st := StopTrying("Welp!").Attach("Max retries attained", 17).Attach("Got this response", "FLOOP").(*internal.AsyncSignalErrorImpl)
				Ω(st.Attachments).Should(HaveLen(2))
				Ω(st.Attachments[0]).Should(Equal(internal.AsyncSignalErrorAttachment{"Max retries attained", 17}))
				Ω(st.Attachments[1]).Should(Equal(internal.AsyncSignalErrorAttachment{"Got this response", "FLOOP"}))
			})
		})

		Describe("when invoking Now()", func() {
			It("should panic with itself", func() {
				st := StopTrying("bam").(*internal.AsyncSignalErrorImpl)
				Ω(st.Now).Should(PanicWith(st))
			})
		})

		Describe("AsAsyncSignalError", func() {
			It("should return false for nils", func() {
				st, ok := internal.AsAsyncSignalError(nil)
				Ω(st).Should(BeNil())
				Ω(ok).Should(BeFalse())
			})

			It("should work when passed a StopTrying error", func() {
				st, ok := internal.AsAsyncSignalError(StopTrying("bam"))
				Ω(st).Should(Equal(StopTrying("bam")))
				Ω(ok).Should(BeTrue())
			})

			It("should work when passed a wrapped error", func() {
				st, ok := internal.AsAsyncSignalError(fmt.Errorf("STOP TRYING %w", StopTrying("bam")))
				Ω(st).Should(Equal(StopTrying("bam")))
				Ω(ok).Should(BeTrue())
			})
		})
	})
})
