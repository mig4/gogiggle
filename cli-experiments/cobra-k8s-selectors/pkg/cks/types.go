package cks

// ListOptions allow configuring a `list` call
type ListOptions struct {
	// LabelSelector specifies a selector to filter the list of returned
	// objects using a K8s-compatible label selector syntax.
	// Defaults to accepting everything.
	LabelSelector string

	// FieldSelector specifies a selector to filter the list of returned
	// objects using a K8s-compatible field selector syntax.
	// Defaults to accepting everything.
	FieldSelector string
}
