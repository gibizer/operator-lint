package a

import (
	. "github.com/onsi/gomega"
)

func DotImportNoLocalGomegaVariable() {
	Eventually(func() {
		Expect(1).To(Equal(1))
	})
}

func DotImportNotUsingLocalGomegaVariable() {
	Eventually(func(g Gomega) {
		g.Expect(1).To(Equal(1))
		Expect(1).To(Equal(1))
		g.Expect(1).To(Equal(1))
	})
}
