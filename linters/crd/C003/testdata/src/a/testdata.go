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
}

type Bad struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	// OptionalField4 is optional by both comment and tag with a default
	OptionalField4 int32 `json:"optionalField4,omitempty"` // want `C003: Field 'OptionalField4' has both a 'Optional' kubebuilder marker with a default value and an 'omitempty' tag. Either remove the default value or remove 'omitempty'`
}
