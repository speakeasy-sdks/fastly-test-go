// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package fastly

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
)

// wafActiveRules - An active rule represents a [rule revision](/reference/api/waf/rules/revisions/) added to a particular [firewall version](/reference/api/waf/firewall-version/).
//
// https://developer.fastly.com/reference/api/waf/rules/active
type wafActiveRules struct {
	sdkConfiguration sdkConfiguration
}

func newWafActiveRules(sdkConfig sdkConfiguration) *wafActiveRules {
	return &wafActiveRules{
		sdkConfiguration: sdkConfig,
	}
}

// BulkUpdateWafActiveRules - Update multiple active rules
// Bulk update all active rules on a [firewall version](https://developer.fastly.com/reference/api/waf/firewall-version/). This endpoint will not add new active rules, only update existing active rules.
//
// Deprecated method: This will be removed in a future release, please migrate away from it as soon as possible.
func (s *wafActiveRules) BulkUpdateWafActiveRules(ctx context.Context, request operations.BulkUpdateWafActiveRulesRequest) (*operations.BulkUpdateWafActiveRulesResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/bulk", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "BulkWafActiveRule", "json", `request:"mediaType=application/vnd.api+json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.BulkUpdateWafActiveRulesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}

// CreateWafActiveRule - Add a rule to a WAF as an active rule
// Create an active rule for a particular firewall version.
//
// Deprecated method: This will be removed in a future release, please migrate away from it as soon as possible.
func (s *wafActiveRules) CreateWafActiveRule(ctx context.Context, request operations.CreateWafActiveRuleRequest) (*operations.CreateWafActiveRuleResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "BulkWafActiveRules", "json", `request:"mediaType=application/vnd.api+json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.CreateWafActiveRuleResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.api+json`):
			var out shared.WafActiveRuleCreationResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.WafActiveRuleCreationResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// CreateWafActiveRulesTag - Create active rules by tag
// Create active rules by tag. This endpoint will create active rules using the latest revision available for each rule.
//
// Deprecated method: This will be removed in a future release, please migrate away from it as soon as possible.
func (s *wafActiveRules) CreateWafActiveRulesTag(ctx context.Context, request operations.CreateWafActiveRulesTagRequest) (*operations.CreateWafActiveRulesTagResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/waf/firewalls/{firewall_id}/versions/{version_id}/tags/{waf_tag_name}/active-rules", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "WafActiveRule1", "json", `request:"mediaType=application/vnd.api+json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.CreateWafActiveRulesTagResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}

// DeleteWafActiveRule - Delete an active rule
// Delete an active rule for a particular firewall version.
//
// Deprecated method: This will be removed in a future release, please migrate away from it as soon as possible.
func (s *wafActiveRules) DeleteWafActiveRule(ctx context.Context, request operations.DeleteWafActiveRuleRequest) (*operations.DeleteWafActiveRuleResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.DeleteWafActiveRuleResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}

// GetWafActiveRule - Get an active WAF rule object
// Get a specific active rule object. Includes details of the rule revision associated with the active rule object by default.
//
// Deprecated method: This will be removed in a future release, please migrate away from it as soon as possible.
func (s *wafActiveRules) GetWafActiveRule(ctx context.Context, request operations.GetWafActiveRuleRequest) (*operations.GetWafActiveRuleResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetWafActiveRuleResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.api+json`):
			var out shared.WafActiveRuleResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.WafActiveRuleResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// ListWafActiveRules - List active rules on a WAF
// List all active rules for a particular firewall version.
//
// Deprecated method: This will be removed in a future release, please migrate away from it as soon as possible.
func (s *wafActiveRules) ListWafActiveRules(ctx context.Context, request operations.ListWafActiveRulesRequest) (*operations.ListWafActiveRulesResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

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

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListWafActiveRulesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.api+json`):
			var out shared.WafActiveRulesResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.WafActiveRulesResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// UpdateWafActiveRule - Update an active rule
// Update an active rule's status for a particular firewall version.
//
// Deprecated method: This will be removed in a future release, please migrate away from it as soon as possible.
func (s *wafActiveRules) UpdateWafActiveRule(ctx context.Context, request operations.UpdateWafActiveRuleRequest) (*operations.UpdateWafActiveRuleResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/waf/firewalls/{firewall_id}/versions/{version_id}/active-rules/{waf_rule_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "WafActiveRule1", "json", `request:"mediaType=application/vnd.api+json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

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

	res := &operations.UpdateWafActiveRuleResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.api+json`):
			var out shared.WafActiveRuleResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.WafActiveRuleResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
