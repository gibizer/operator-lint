# T001

T001 checks that Gomega's `Eventually` and `Consistently` blocks use a local
Gomega instance for asserts instead of a global one.

It is a common pattern to make multiple assertions on a single
object in an `Eventually` or `Consistently` block:

```golang
func ExpectCondition(name string, expectedStatus string, expectedReason string) {
    Eventually(func(g Gomega) {
		conditions := getter.GetConditions(name)
		g.Expect(conditions).NotTo(BeNil())
		g.Expect(actualCondition.Status).To(Equal(expectedStatus))
        // This can fail silently
		Expect(actualCondition.Reason).To(Equal(expectedReason))
	}, timeout, interval).Should(Succeed())
}
```

However if the code calls the global `Expect` function from the `gomega` module
instead of using the locally passed in `g Gomega` variable then such `Except`
call can fail silently without properly stopping and cleaning up the test. This
is documented in the
[gomega doc](https://onsi.github.io/gomega/#category-3-making-assertions-eminem-the-function-passed-into-codeeventuallycode)
but nothing is enforcing this rule in `gomega`.
