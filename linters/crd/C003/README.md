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

Similary this issue can happen when a defaulting webhook deserializing the json
request, run the specific defaulting code and then serializing the struct back
to json.

This behavior does not cause an issue if the field's default value is the same
as the golang empty value of the type of the field.

If `omitempty` cannot be removed from the field definition (e.g. nested struct
defaulting), then one possible way out is to switch the field type to be a
pointer. E.g. `Replicas *int32` in the above example. This introduces a new
empty value to the field `nil`. So if the user never wants to explicitly set
`nil`, then it is a good value to represent the `unset` state via `omitempty`.

The C003 check will not report an error for pointer type fields or field
definitions where the kubebuilder default value is the same as the golang
empty value of the type of the field.
