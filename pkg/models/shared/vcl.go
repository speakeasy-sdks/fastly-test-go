// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Vcl struct {
	// The VCL code to be included.
	Content *string `form:"name=content"`
	// Set to `true` when this is the main VCL, otherwise `false`.
	Main *bool `form:"name=main"`
	// The name of this VCL.
	Name *string `form:"name=name"`
}
