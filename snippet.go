// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package fastlytestgo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/hooks"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/operations"
	"github.com/speakeasy-sdks/fastly-test-go/v2/models/sdkerrors"
	"io"
	"net/http"
)

// Snippet - VCL Snippets are blocks of VCL logic inserted into your service's configuration that don't require custom VCL.
//
// https://developer.fastly.com/reference/api/vcl-services/snippet
type Snippet struct {
	sdkConfiguration sdkConfiguration
}

func newSnippet(sdkConfig sdkConfiguration) *Snippet {
	return &Snippet{
		sdkConfiguration: sdkConfig,
	}
}

// CreateSnippet - Create a snippet
// Create a snippet for a particular service and version.
func (s *Snippet) CreateSnippet(ctx context.Context, request operations.CreateSnippetRequest) (*operations.CreateSnippetResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "create-snippet"}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/service/{service_id}/version/{version_id}/snippet", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Snippet", "form", `request:"mediaType=application/x-www-form-urlencoded"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateSnippetResponse{
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
		case utils.MatchContentType(contentType, `application/json`):
			var out components.SnippetResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.SnippetResponse = &out
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

// DeleteSnippet - Delete a snippet
// Delete a specific snippet for a particular service and version.
func (s *Snippet) DeleteSnippet(ctx context.Context, request operations.DeleteSnippetRequest) (*operations.DeleteSnippetResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "delete-snippet"}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/service/{service_id}/version/{version_id}/snippet/{snippet_name}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteSnippetResponse{
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
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.DeleteSnippetResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
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

// GetSnippet - Get a versioned snippet
// Get a single snippet for a particular service and version.
func (s *Snippet) GetSnippet(ctx context.Context, request operations.GetSnippetRequest) (*operations.GetSnippetResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "get-snippet"}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/service/{service_id}/version/{version_id}/snippet/{snippet_name}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSnippetResponse{
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
		case utils.MatchContentType(contentType, `application/json`):
			var out components.SnippetResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.SnippetResponse = &out
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

// GetSnippetDynamic - Get a dynamic snippet
// Get a single dynamic snippet for a particular service.
func (s *Snippet) GetSnippetDynamic(ctx context.Context, request operations.GetSnippetDynamicRequest) (*operations.GetSnippetDynamicResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "get-snippet-dynamic"}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/service/{service_id}/snippet/{snippet_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSnippetDynamicResponse{
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
		case utils.MatchContentType(contentType, `application/json`):
			var out components.SnippetResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.SnippetResponse = &out
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

// ListSnippets - List snippets
// List all snippets for a particular service and version.
func (s *Snippet) ListSnippets(ctx context.Context, request operations.ListSnippetsRequest) (*operations.ListSnippetsResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "list-snippets"}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/service/{service_id}/version/{version_id}/snippet", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListSnippetsResponse{
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
		case utils.MatchContentType(contentType, `application/json`):
			var out []components.SnippetResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Classes = out
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

// UpdateSnippetDynamic - Update a dynamic snippet
// Update a dynamic snippet for a particular service.
func (s *Snippet) UpdateSnippetDynamic(ctx context.Context, request operations.UpdateSnippetDynamicRequest) (*operations.UpdateSnippetDynamicResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "update-snippet-dynamic"}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/service/{service_id}/snippet/{snippet_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Snippet", "form", `request:"mediaType=application/x-www-form-urlencoded"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateSnippetDynamicResponse{
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
		case utils.MatchContentType(contentType, `application/json`):
			var out components.SnippetResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.SnippetResponse = &out
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
