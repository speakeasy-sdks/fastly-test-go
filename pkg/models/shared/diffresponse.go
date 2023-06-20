// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DiffResponse - OK
type DiffResponse struct {
	// The differences between two specified service versions. Returns the full config if the version configurations are identical.
	Diff *string `json:"diff,omitempty"`
	// The format the diff is being returned in (`text`, `html` or `html_simple`).
	Format *string `json:"format,omitempty"`
	// The version number being diffed from.
	From *int64 `json:"from,omitempty"`
	// The version number being diffed to.
	To *int64 `json:"to,omitempty"`
}