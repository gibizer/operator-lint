# C001

C001 detects incompatible `Required` and `Optional` kubebuilder markers.

```golang
// +kubebuilder:Optional
// +kubebuilder:Required
Foo string `json:"foo"`
```
It is undefined that the field `Foo` will be required or optional.
