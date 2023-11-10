// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
)

// HTTPResponseFormat - Payload format for delivering to subscribers of whole HTTP responses (`response` hold mode). One of `body` or `body-bin` must be specified.
type HTTPResponseFormat struct {
	// The response body as a string.
	Body *string `json:"body,omitempty"`
	// The response body as a base64-encoded binary blob.
	BodyBin *string `json:"body-bin,omitempty"`
	// The HTTP status code.
	Code *int64 `default:"200" json:"code"`
	// A map of arbitrary HTTP response headers to include in the response.
	Headers map[string]string `json:"headers,omitempty"`
	// The HTTP status string. Defaults to a string appropriate for `code`.
	Reason *string `json:"reason,omitempty"`
}

func (h HTTPResponseFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HTTPResponseFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HTTPResponseFormat) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *HTTPResponseFormat) GetBodyBin() *string {
	if o == nil {
		return nil
	}
	return o.BodyBin
}

func (o *HTTPResponseFormat) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *HTTPResponseFormat) GetHeaders() map[string]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *HTTPResponseFormat) GetReason() *string {
	if o == nil {
		return nil
	}
	return o.Reason
}
