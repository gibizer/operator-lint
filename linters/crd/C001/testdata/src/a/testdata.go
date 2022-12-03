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

	NotCommentedField string
}

type Bad struct {
	//+kubebuilder:validation:Required
	// +kubebuilder:validation:Optional
	ContradictingMarkerField string // want `C001: Field 'ContradictingMarkerField' has both 'Optional' and 'Required' kubebuilder markers. Only one of them should be used`

	//+kubebuilder:validation:Optional
	// ContradictingMarkerField2 -
	// +kubebuilder:validation:Required
	ContradictingMarkerField2 string // want `C001: Field 'ContradictingMarkerField2' has both 'Optional' and 'Required' kubebuilder markers. Only one of them should be used`
}
