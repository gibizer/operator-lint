package a

import (
	gomega "github.com/onsi/gomega"
)

func NoLocalGomegaVariable() {
	gomega.Eventually(func() {
		gomega.Expect(1).To(gomega.Equal(1))
	})
}

func NotUsingLocalGomegaVariable() {
	gomega.Eventually(func(g gomega.Gomega) {
		g.Expect(1).To(gomega.Equal(1))
		gomega.Expect(1).To(gomega.Equal(1))
		g.Expect(1).To(gomega.Equal(1))
	})
}
