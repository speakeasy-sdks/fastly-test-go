// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"Fastly/pkg/models/operations"
	"Fastly/pkg/models/sdkerrors"
	"Fastly/pkg/models/shared"
	"Fastly/pkg/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// tlsCertificates - A TLS certificate is used to terminate TLS traffic for one or more of your [TLS domains](#tls_domains).
// https://developer.fastly.com/reference/api/tls/custom-certs/certificates
type tlsCertificates struct {
	sdkConfiguration sdkConfiguration
}

func newTLSCertificates(sdkConfig sdkConfiguration) *tlsCertificates {
	return &tlsCertificates{
		sdkConfiguration: sdkConfig,
	}
}

// CreateTLSCert - Create a TLS certificate
// Create a TLS certificate.
func (s *tlsCertificates) CreateTLSCert(ctx context.Context, request shared.TLSCertificateInput, security operations.CreateTLSCertSecurity) (*operations.CreateTLSCertResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/tls/certificates"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

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

	res := &operations.CreateTLSCertResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.api+json`):
			var out *operations.CreateTLSCert201ApplicationVndAPIPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.CreateTLSCert201ApplicationVndAPIPlusJSONObject = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// DeleteTLSCert - Delete a TLS certificate
// Destroy a TLS certificate. TLS certificates already enabled for a domain cannot be destroyed.
func (s *tlsCertificates) DeleteTLSCert(ctx context.Context, request operations.DeleteTLSCertRequest, security operations.DeleteTLSCertSecurity) (*operations.DeleteTLSCertResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tls/certificates/{tls_certificate_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

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

	res := &operations.DeleteTLSCertResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}

// GetTLSCert - Get a TLS certificate
// Show a TLS certificate.
func (s *tlsCertificates) GetTLSCert(ctx context.Context, request operations.GetTLSCertRequest, security operations.GetTLSCertSecurity) (*operations.GetTLSCertResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tls/certificates/{tls_certificate_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

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

	res := &operations.GetTLSCertResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.api+json`):
			var out *shared.TLSCertificateResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.TLSCertificateResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// ListTLSCerts - List TLS certificates
// List all TLS certificates.
func (s *tlsCertificates) ListTLSCerts(ctx context.Context, request operations.ListTLSCertsRequest, security operations.ListTLSCertsSecurity) (*operations.ListTLSCertsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/tls/certificates"

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

	res := &operations.ListTLSCertsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.api+json`):
			var out *shared.TLSCertificatesResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.TLSCertificatesResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// UpdateTLSCert - Update a TLS certificate
// Replace a TLS certificate with a newly reissued TLS certificate, or update a TLS certificate's name. If replacing a TLS certificate, the new TLS certificate must contain all SAN entries as the current TLS certificate. It must either have an exact matching list or contain a superset.
func (s *tlsCertificates) UpdateTLSCert(ctx context.Context, request operations.UpdateTLSCertRequest, security operations.UpdateTLSCertSecurity) (*operations.UpdateTLSCertResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/tls/certificates/{tls_certificate_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "TLSCertificateInput", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

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

	res := &operations.UpdateTLSCertResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.api+json`):
			var out *shared.TLSCertificateResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.TLSCertificateResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
