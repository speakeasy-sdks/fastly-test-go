// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"net/http"
)

type ListTLSDomainsRequest struct {
	// Optional. Limit the returned domains to those currently using Fastly to terminate TLS with SNI (that is, domains considered "in use") Permitted values: true, false.
	FilterInUse *string `queryParam:"style=form,explode=true,name=filter[in_use]"`
	// Optional. Limit the returned domains to those listed in the given TLS certificate's SAN list.
	FilterTLSCertificatesID *string `queryParam:"style=form,explode=true,name=filter[tls_certificates.id]"`
	// Optional. Limit the returned domains to those for a given TLS subscription.
	FilterTLSSubscriptionsID *string `queryParam:"style=form,explode=true,name=filter[tls_subscriptions.id]"`
	// Include related objects. Optional, comma-separated values. Permitted values: `tls_activations`, `tls_certificates`, `tls_subscriptions`, `tls_subscriptions.tls_authorizations`, and `tls_authorizations.globalsign_email_challenge`.
	//
	Include *string `queryParam:"style=form,explode=true,name=include"`
	// Current page.
	PageNumber *int64 `queryParam:"style=form,explode=true,name=page[number]"`
	// Number of records per page.
	PageSize *int64 `default:"20" queryParam:"style=form,explode=true,name=page[size]"`
	// The order in which to list the results by creation date.
	Sort *components.Sort `default:"created_at" queryParam:"style=form,explode=true,name=sort"`
}

func (l ListTLSDomainsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTLSDomainsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTLSDomainsRequest) GetFilterInUse() *string {
	if o == nil {
		return nil
	}
	return o.FilterInUse
}

func (o *ListTLSDomainsRequest) GetFilterTLSCertificatesID() *string {
	if o == nil {
		return nil
	}
	return o.FilterTLSCertificatesID
}

func (o *ListTLSDomainsRequest) GetFilterTLSSubscriptionsID() *string {
	if o == nil {
		return nil
	}
	return o.FilterTLSSubscriptionsID
}

func (o *ListTLSDomainsRequest) GetInclude() *string {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *ListTLSDomainsRequest) GetPageNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.PageNumber
}

func (o *ListTLSDomainsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListTLSDomainsRequest) GetSort() *components.Sort {
	if o == nil {
		return nil
	}
	return o.Sort
}

type ListTLSDomainsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TLSDomainsResponse *components.TLSDomainsResponse
}

func (o *ListTLSDomainsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTLSDomainsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTLSDomainsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTLSDomainsResponse) GetTLSDomainsResponse() *components.TLSDomainsResponse {
	if o == nil {
		return nil
	}
	return o.TLSDomainsResponse
}
