package a

import (
	gg "github.com/onsi/gomega"
)

func NamedImportNoLocalGomegaVariable() {
	gg.Eventually(func() {
		gg.Expect(1).To(gg.Equal(1)) // want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Change the declaration of the function passed to the 'Eventually' block to take a parameter with type 'Gomega' and use that to call 'Expect'`
	})
}

func NamedImportNotUsingLocalGomegaVariable() {
	gg.Eventually(func(g gg.Gomega) {
		g.Expect(1).To(gg.Equal(1))
		gg.Expect(1).To(gg.Equal(1)) // want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Use 'g.Expect'`
		g.Expect(1).To(gg.Equal(1))
	})
}

func NamedImportPredefinedFunction() {
	f := func() {
		gg.Expect(1).To(gg.Equal(1)) // TODO: want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Change the declaration of the function passed to the 'Eventually' block to take a parameter with type 'Gomega' and use that to call 'Expect'`
	}
	gg.Eventually(f)
}

func NamedImportPredefinedFunctionUnusedLocal() {
	f := func(g gg.Gomega) {
		gg.Expect(1).To(gg.Equal(1)) // TODO: want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Use 'g.Expect'`
	}
	gg.Eventually(f)
}
