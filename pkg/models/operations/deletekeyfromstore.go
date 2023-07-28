// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteKeyFromStoreSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *DeleteKeyFromStoreSecurity) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

type DeleteKeyFromStoreRequest struct {
	Force   *bool  `queryParam:"style=form,explode=true,name=force"`
	KeyName string `pathParam:"style=simple,explode=false,name=key_name"`
	StoreID string `pathParam:"style=simple,explode=false,name=store_id"`
}

func (o *DeleteKeyFromStoreRequest) GetForce() *bool {
	if o == nil {
		return nil
	}
	return o.Force
}

func (o *DeleteKeyFromStoreRequest) GetKeyName() string {
	if o == nil {
		return ""
	}
	return o.KeyName
}

func (o *DeleteKeyFromStoreRequest) GetStoreID() string {
	if o == nil {
		return ""
	}
	return o.StoreID
}

type DeleteKeyFromStoreResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *DeleteKeyFromStoreResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteKeyFromStoreResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteKeyFromStoreResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
