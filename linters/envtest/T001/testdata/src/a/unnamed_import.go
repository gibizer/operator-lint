package a

import (
	"github.com/onsi/gomega"
)

func UnnamedImportNoLocalGomegaVariable() {
	gomega.Eventually(func() {
		gomega.Expect(1).To(gomega.Equal(1)) // want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Change the declaration of the function passed to the 'Eventually' block to take a parameter with type 'Gomega' and use that to call 'Expect'`
	})
}

func UnnamedImportNotUsingLocalGomegaVariable() {
	gomega.Eventually(func(g gomega.Gomega) {
		g.Expect(1).To(gomega.Equal(1))
		gomega.Expect(1).To(gomega.Equal(1)) // want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Use 'g.Expect'`
		g.Expect(1).To(gomega.Equal(1))
	})
}

func UnnamedImportPredefinedFunction() {
	f := func() {
		gomega.Expect(1).To(gomega.Equal(1)) // TODO: want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Change the declaration of the function passed to the 'Eventually' block to take a parameter with type 'Gomega' and use that to call 'Expect'`
	}
	gomega.Eventually(f)
}

func UnnamedImportPredefinedFunctionUnusedLocal() {
	f := func(g gomega.Gomega) {
		gomega.Expect(1).To(gomega.Equal(1)) // TODO: want `L001: 'Expect' in the 'Eventually' block should be called via a local 'Gomega' parameter. Use 'g.Expect'`
	}
	gomega.Eventually(f)
}
