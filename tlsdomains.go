// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package fastlytestgo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/sdkerrors"
	"io"
	"net/http"
	"strings"
)

// TLSDomains - TLS domains are all the domains (including wildcard domains) included in any [TLS certificate](#tls_certificates)'s Subject Alternative Names (SAN) list. Included in the response is information about which certificates reference this domain as well as the [TLS activation](#tls_activations) indicating which certificate is enabled to serve TLS traffic for the domain.
//
// https://developer.fastly.com/reference/api/tls/custom-certs/domains
type TLSDomains struct {
	sdkConfiguration sdkConfiguration
}

func newTLSDomains(sdkConfig sdkConfiguration) *TLSDomains {
	return &TLSDomains{
		sdkConfiguration: sdkConfig,
	}
}

// ListTLSDomains - List TLS domains
// List all TLS domains.
func (s *TLSDomains) ListTLSDomains(ctx context.Context, request operations.ListTLSDomainsRequest) (*operations.ListTLSDomainsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/tls/domains"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListTLSDomainsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.api+json`):
			var out components.TLSDomainsResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TLSDomainsResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
