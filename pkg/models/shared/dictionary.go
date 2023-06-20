// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Dictionary struct {
	// Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace).
	Name *string `form:"name=name"`
	// Determines if items in the dictionary are readable or not.
	WriteOnly *bool `form:"name=write_only"`
}