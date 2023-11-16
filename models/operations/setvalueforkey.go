// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type SetValueForKeyRequest struct {
	RequestBody       *string `request:"mediaType=application/octet-stream"`
	Add               *bool   `queryParam:"style=form,explode=true,name=add"`
	Append            *bool   `queryParam:"style=form,explode=true,name=append"`
	BackgroundFetch   *bool   `queryParam:"style=form,explode=true,name=background_fetch"`
	IfGenerationMatch *int64  `header:"style=simple,explode=false,name=if-generation-match"`
	KeyName           string  `pathParam:"style=simple,explode=false,name=key_name"`
	Metadata          *string `header:"style=simple,explode=false,name=metadata"`
	Prepend           *bool   `queryParam:"style=form,explode=true,name=prepend"`
	StoreID           string  `pathParam:"style=simple,explode=false,name=store_id"`
	TimeToLiveSec     *int64  `header:"style=simple,explode=false,name=time_to_live_sec"`
}

func (o *SetValueForKeyRequest) GetRequestBody() *string {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *SetValueForKeyRequest) GetAdd() *bool {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *SetValueForKeyRequest) GetAppend() *bool {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *SetValueForKeyRequest) GetBackgroundFetch() *bool {
	if o == nil {
		return nil
	}
	return o.BackgroundFetch
}

func (o *SetValueForKeyRequest) GetIfGenerationMatch() *int64 {
	if o == nil {
		return nil
	}
	return o.IfGenerationMatch
}

func (o *SetValueForKeyRequest) GetKeyName() string {
	if o == nil {
		return ""
	}
	return o.KeyName
}

func (o *SetValueForKeyRequest) GetMetadata() *string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *SetValueForKeyRequest) GetPrepend() *bool {
	if o == nil {
		return nil
	}
	return o.Prepend
}

func (o *SetValueForKeyRequest) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

func (o *SetValueForKeyRequest) GetTimeToLiveSec() *int64 {
	if o == nil {
		return nil
	}
	return o.TimeToLiveSec
}

type SetValueForKeyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Res *string
}

func (o *SetValueForKeyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetValueForKeyResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *SetValueForKeyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetValueForKeyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SetValueForKeyResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
