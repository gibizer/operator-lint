package a

import (
	. "github.com/onsi/gomega"
)

func DotImportNoLocalGomegaVariable() {
	Eventually(func() {
		Expect(1).To(Equal(1)) // want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Change the declaration of the function passed to the 'Eventually' block to take a parameter with type 'Gomega' and use that to call 'Expect'`
	})
}

func DotImportNotUsingLocalGomegaVariable() {
	Eventually(func(g Gomega) {
		g.Expect(1).To(Equal(1))
		Expect(1).To(Equal(1)) // want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Use 'g.Expect'`
		g.Expect(1).To(Equal(1))
	})
}

func DotImportPredefinedFunction() {
	f := func() {
		Expect(1).To(Equal(1)) // TODO: want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Change the declaration of the function passed to the 'Eventually' block to take a parameter with type 'Gomega' and use that to call 'Expect'`
	}
	Eventually(f)
}

func DotImportPredefinedFunctionUnusedLocal() {
	f := func(g Gomega) {
		Expect(1).To(Equal(1)) // TODO: want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Use 'g.Expect'`
	}
	Eventually(f)
}
