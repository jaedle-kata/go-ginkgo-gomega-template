package kata_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jaedle-kata/go-ginkgo-gomega-template/kata"
)

var _ = Describe("Kata", func() {
	It("runs", func() {
		Expect(kata.Kata()).To(Equal(true))
	})
})
