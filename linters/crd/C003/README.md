# C003

C003 detects incompatible defaulting via `Optional` kubebuilder marker and
`omitemty` golang tag

The `omitempty` tag can cause surprising behavior when used on a field that has
default value defined. For example:
If we have field
```golang
// +kubebuilder:default=1
Replicas int32 `json:"replicas,omitempty"`
```
and the client is populating this field via `controllerutil.CreateOrPatch` with
a value `0` then the value of the field will be `1` instead of `0` on the
server side. The `0` value is the *empty* value of `int32` and `omitempty`
means if the value is empty then the field is not emitted when serialized to
json. Therefore the server won't get the field `replicas` and therefore uses
the default value `1` from the kubebuilder default annotation.
