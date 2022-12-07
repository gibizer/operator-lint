# C002

C002 detects incompatible `Required` kubebuilder marker and `omitemty` golang tag.

```golang
// +kubebuilder:Required
Foo string `json:"foo,omitempty"`
```

The `omitempty` tag makes the field `Foo` optional regardless of the `Required`
marker so this makes the code confusing.
