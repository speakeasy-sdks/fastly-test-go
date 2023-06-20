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
)

// enabledProducts - These endpoints allow you to enable, disable, and check the enablement status of products on your services.
// https://developer.fastly.com/reference/api/products/enablement
type enabledProducts struct {
	sdkConfiguration sdkConfiguration
}

func newEnabledProducts(sdkConfig sdkConfiguration) *enabledProducts {
	return &enabledProducts{
		sdkConfiguration: sdkConfig,
	}
}

// DisableProduct - Disable a product
// Disable a product on a service. Supported product IDs: `brotli_compression`,`domain_inspector`,`fanout`,`image_optimizer`,`origin_inspector`, and `websockets`.
func (s *enabledProducts) DisableProduct(ctx context.Context, request operations.DisableProductRequest, security operations.DisableProductSecurity) (*operations.DisableProductResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/enabled-products/{product_id}/services/{service_id}", request, nil)
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

	res := &operations.DisableProductResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}

// EnableProduct - Enable a product
// Enable a product on a service. Supported product IDs: `brotli_compression`,`domain_inspector`,`fanout`,`image_optimizer`,`origin_inspector`, and `websockets`.
func (s *enabledProducts) EnableProduct(ctx context.Context, request operations.EnableProductRequest, security operations.EnableProductSecurity) (*operations.EnableProductResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/enabled-products/{product_id}/services/{service_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
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

	res := &operations.EnableProductResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.EnabledProductResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.EnabledProductResponse = out
		}
	}

	return res, nil
}

// GetEnabledProduct - Get enabled product
// Get enabled product on a service. Supported product IDs: `brotli_compression`,`domain_inspector`,`fanout`,`image_optimizer`,`origin_inspector`, and `websockets`.
func (s *enabledProducts) GetEnabledProduct(ctx context.Context, request operations.GetEnabledProductRequest, security operations.GetEnabledProductSecurity) (*operations.GetEnabledProductResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/enabled-products/{product_id}/services/{service_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
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

	res := &operations.GetEnabledProductResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.EnabledProductResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.EnabledProductResponse = out
		}
	}

	return res, nil
}
