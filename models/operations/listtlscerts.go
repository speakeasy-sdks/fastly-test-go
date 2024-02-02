// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type ListTLSCertsRequest struct {
	// Optional. Limit the returned certificates to those currently using Fastly to terminate TLS (that is, certificates associated with an activation). Permitted values: true, false.
	FilterInUse *string `queryParam:"style=form,explode=true,name=filter[in_use]"`
	// Limit the returned certificates to those that expire prior to the specified date in UTC. Accepts parameters: lte (e.g., filter[not_after][lte]=2020-05-05).
	//
	FilterNotAfter *string `queryParam:"style=form,explode=true,name=filter[not_after]"`
	// Limit the returned certificates to those that include the specific domain.
	FilterTLSDomainsID *string `queryParam:"style=form,explode=true,name=filter[tls_domains.id]"`
	// Include related objects. Optional, comma-separated values. Permitted values: `tls_activations`.
	//
	Include *string `queryParam:"style=form,explode=true,name=include"`
	// Current page.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Number of records per page.
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=page[size]"`
	// The order in which to list the results by creation date.
	Sort *components.Sort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
}

func (l ListTLSCertsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTLSCertsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTLSCertsRequest) GetFilterInUse() *string {
	if o == nil {
		return nil
	}
	return o.FilterInUse
}

func (o *ListTLSCertsRequest) GetFilterNotAfter() *string {
	if o == nil {
		return nil
	}
	return o.FilterNotAfter
}

func (o *ListTLSCertsRequest) GetFilterTLSDomainsID() *string {
	if o == nil {
		return nil
	}
	return o.FilterTLSDomainsID
}

func (o *ListTLSCertsRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *ListTLSCertsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListTLSCertsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListTLSCertsRequest) GetSort() *components.Sort {
	if o == nil {
		return nil
	}
	return o.Sort
}

type ListTLSCertsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TLSCertificatesResponse *components.TLSCertificatesResponse
}

func (o *ListTLSCertsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTLSCertsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTLSCertsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTLSCertsResponse) GetTLSCertificatesResponse() *components.TLSCertificatesResponse {
	if o == nil {
		return nil
	}
	return o.TLSCertificatesResponse
}
