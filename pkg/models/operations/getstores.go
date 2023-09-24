// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"Fastly/pkg/utils"
	"net/http"
)

type GetStoresRequest struct {
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	Limit  *int64  `default:"100" queryParam:"style=form,explode=true,name=limit"`
}

func (g GetStoresRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetStoresRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetStoresRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *GetStoresRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

// GetStores200ApplicationJSONMeta - Meta for the pagination.
type GetStores200ApplicationJSONMeta struct {
	// Entries returned.
	Limit *int64 `json:"limit,omitempty"`
	// Cursor for the next page.
	NextCursor *string `json:"next_cursor,omitempty"`
}

func (o *GetStores200ApplicationJSONMeta) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetStores200ApplicationJSONMeta) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

// GetStores200ApplicationJSON - OK
type GetStores200ApplicationJSON struct {
	Data []shared.StoreResponse `json:"data,omitempty"`
	// Meta for the pagination.
	Meta *GetStores200ApplicationJSONMeta `json:"meta,omitempty"`
}

func (o *GetStores200ApplicationJSON) GetData() []shared.StoreResponse {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetStores200ApplicationJSON) GetMeta() *GetStores200ApplicationJSONMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

type GetStoresResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	GetStores200ApplicationJSONObject *GetStores200ApplicationJSON
}

func (o *GetStoresResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetStoresResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetStoresResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetStoresResponse) GetGetStores200ApplicationJSONObject() *GetStores200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetStores200ApplicationJSONObject
}
