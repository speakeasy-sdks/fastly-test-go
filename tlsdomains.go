// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// tlsDomains - TLS domains are all the domains (including wildcard domains) included in any [TLS certificate](#tls_certificates)'s Subject Alternative Names (SAN) list. Included in the response is information about which certificates reference this domain as well as the [TLS activation](#tls_activations) indicating which certificate is enabled to serve TLS traffic for the domain.
// https://developer.fastly.com/reference/api/tls/custom-certs/domains
type tlsDomains struct {
	sdkConfiguration sdkConfiguration
}

func newTLSDomains(sdkConfig sdkConfiguration) *tlsDomains {
	return &tlsDomains{
		sdkConfiguration: sdkConfig,
	}
}

// ListTLSDomains - List TLS domains
// List all TLS domains.
func (s *tlsDomains) ListTLSDomains(ctx context.Context, request operations.ListTLSDomainsRequest, security operations.ListTLSDomainsSecurity) (*operations.ListTLSDomainsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/tls/domains"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListTLSDomainsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.api+json`):
			var out *shared.TLSDomainsResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.TLSDomainsResponse = out
		}
	}

	return res, nil
}
