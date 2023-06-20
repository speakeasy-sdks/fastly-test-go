// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type VersionInput struct {
	// Whether this is the active version or not.
	Active *bool `form:"name=active"`
	// A freeform descriptive note.
	Comment *string `form:"name=comment"`
	// Unused at this time.
	Deployed *bool `form:"name=deployed"`
	// Whether this version is locked or not. Objects can not be added or edited on locked versions.
	Locked *bool `form:"name=locked"`
	// Unused at this time.
	Staging *bool `form:"name=staging"`
	// Unused at this time.
	Testing *bool `form:"name=testing"`
}

// Version - OK
type Version struct {
	// Whether this is the active version or not.
	Active *bool `json:"active,omitempty"`
	// A freeform descriptive note.
	Comment *string `json:"comment,omitempty"`
	// Unused at this time.
	Deployed *bool `json:"deployed,omitempty"`
	// Whether this version is locked or not. Objects can not be added or edited on locked versions.
	Locked *bool `json:"locked,omitempty"`
	// The number of this version.
	Number *int64 `json:"number,omitempty"`
	// Unused at this time.
	Staging *bool `json:"staging,omitempty"`
	// Unused at this time.
	Testing *bool `json:"testing,omitempty"`
}
