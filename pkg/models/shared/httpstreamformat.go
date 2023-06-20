// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// HTTPStreamFormat - Payload format for delivering to subscribers of HTTP streaming response bodies (`stream` hold mode). One of `content` or `content-bin` must be specified.
type HTTPStreamFormat struct {
	// A fragment of body data as a string.
	Content *string `json:"content,omitempty"`
	// A fragment of body data as a base64-encoded binary blob.
	ContentBin *string `json:"content-bin,omitempty"`
}
