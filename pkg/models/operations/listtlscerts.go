// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type ListTLSCertsSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

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
	PageSize *int64 `queryParam:"style=form,explode=true,name=page[size]"`
	// The order in which to list the results by creation date.
	Sort *shared.Sort `queryParam:"style=form,explode=true,name=sort"`
}

type ListTLSCertsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	TLSCertificatesResponse *shared.TLSCertificatesResponse
}
