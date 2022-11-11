package a

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Multipe valid cases where no assert happens inside the passed in
// function or no function passed
// https://onsi.github.io/gomega/#eventually
func Valid() {

	Eventually(1).Should(Equal(1))

	fetcher := func() int { return 1 }
	Eventually(fetcher).Should(Equal(1))

	Eventually(func() int {
		return 1
	}).Should(Equal(1))

	It("foo", func(ctx SpecContext) {
		Eventually(ctx, fetcher).Should(Equal(1))
		Eventually(fetcher, ctx).Should(Equal(1))
		Eventually(ctx, func() int {
			return 1
		}).Should(Equal(1))
		Eventually(func() int {
			return 1
		}, ctx).Should(Equal(1))

	})

}
