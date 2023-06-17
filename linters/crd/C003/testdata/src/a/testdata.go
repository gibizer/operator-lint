package a

type Good struct {
	// +kubebuilder:validation:Optional
	// OptionalField is optional by comment
	OptionalField string `json:"optionalField"`

	// OptionalField is optional by tag
	OptionalField2 string `json:"optionalField2,omitempty"`

	NotCommentedField string

	NotCommentedField2 string `json:"notCommentedField2"`

	NotCommentedField3 string `json:"notCommentedField3,omitempty"`

	// OptionalField3 is optional by both comment and tag but no default
	// +kubebuilder:validation:Optional
	OptionalField3 string `json:"optionalField3,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=0
	// OptionalField4 has omitempty and default value as well but the
	// default value is the same as the golang default.
	OptionalField4 int32 `json:"optionalField4,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=0
	// OptionalField5 has omitempty and default value as well but the
	// default value is the same as the golang default.
	OptionalField5 int `json:"optionalField5,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	// OptionalField6 has omitempty and default value as well but the
	// default value is the same as the golang default.
	OptionalField6 bool `json:"optionalField6,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	// OptionalField7 is optional by both comment and tag with a default, the
	// default is not the golang default but the field is a pointer
	OptionalField7 *int32 `json:"optionalField7,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=3
	// OptionalField8 is optional by both comment and tag with a default, the
	// default is not the golang default but the field is a pointer
	OptionalField8 *int `json:"optionalField8,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	// OptionalField9 is optional by both comment and tag with a default, the
	// default is not the golang default but the field is a pointer
	OptionalField9 *bool `json:"optionalField9,omitempty"`
}

type Bad struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	// OptionalField4 is optional by both comment and tag with a default
	OptionalField4 int32 `json:"optionalField4,omitempty"` // want `C003: Field 'OptionalField4' has both a 'Optional' kubebuilder marker with a default value and an 'omitempty' tag. Either remove the default value or remove 'omitempty'`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=3
	// OptionalField5 is optional by both comment and tag with a default
	OptionalField5 int `json:"optionalField5,omitempty"` // want `C003: Field 'OptionalField5' has both a 'Optional' kubebuilder marker with a default value and an 'omitempty' tag. Either remove the default value or remove 'omitempty'`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	// OptionalField6 is optional by both comment and tag with a default
	OptionalField6 bool `json:"optionalField6,omitempty"` // want `C003: Field 'OptionalField6' has both a 'Optional' kubebuilder marker with a default value and an 'omitempty' tag. Either remove the default value or remove 'omitempty'`
}
