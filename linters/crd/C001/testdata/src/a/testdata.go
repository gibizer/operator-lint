package a

type Good struct {
	// +kubebuilder:validation:Required
	// MandatoryField is mandatory
	MandatoryField string `json:"mandatoryField"`

	// MandatoryField2 is mandatory
	// +kubebuilder:validation:Required
	MandatoryField2 string `json:"mandatory2Field"`

	// +kubebuilder:validation:Optional
	// OptionalField is optional by both comment and tag
	OptionalField string `json:"optionalField,omitempty"`

	// OptionalField2 is optional by both comment and tag
	// +kubebuilder:validation:Optional
	OptionalField2 string `json:"optionalField2,omitempty"`

	// +kubebuilder:validation:Optional
	// OptionalField3 is optional by comment
	OptionalField3 string `json:"optionalField3"`

	// OptionalField4 is optional by tag
	OptionalField4 string `json:"optionalField4,omitempty"`

	//+kubebuilder:validation:Optional
	// OptionalField5 is optional by comment
	OptionalField5 string `json:"optionalField5"`

	NotCommentedField string

	NotCommentedField2 string `json:"notCommentedField2"`

	NotCommentedField3 string `json:"notCommentedField3,omitempty"`
}

type Bad struct {
	// +kubebuilder:validation:Required
	// MandatoryField wants to be mandatory but omitempty makes it optional
	MandatoryField string `json:"mandatoryField,omitempty"` // want `C001: Field 'MandatoryField' has both a 'Required' kubebuilder marker and an 'omitempty' tag. Only one of them should be used`

	// MandatoryField2 wants to be mandatory but omitempty makes it optional
	// +kubebuilder:validation:Required
	MandatoryField2 string `json:"mandatoryField2,omitempty"` // want `C001: Field 'MandatoryField2' has both a 'Required' kubebuilder marker and an 'omitempty' tag. Only one of them should be used`

	//+kubebuilder:validation:Required
	// MandatoryField3 wants to be mandatory but omitempty makes it optional
	MandatoryField3 string `json:"mandatoryField3,omitempty"` // want `C001: Field 'MandatoryField3' has both a 'Required' kubebuilder marker and an 'omitempty' tag. Only one of them should be used`

	//+kubebuilder:validation:Required
	// +kubebuilder:validation:Optional
	ContradictingMarkerField string // want `C001: Field 'ContradictingMarkerField' has both 'Optional' and 'Required' kubebuilder markers. Only one of them should be used`
}
