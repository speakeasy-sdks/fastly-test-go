// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package fastlytestgo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"github.com/speakeasy-sdks/fastly-test-go/models/operations"
	"github.com/speakeasy-sdks/fastly-test-go/models/sdkerrors"
	"io"
	"net/http"
)

// Stats give you information on the usage and performance of your Service. They can be requested by Service and over a particular time span. Stats are broken down per POP, giving you information on how your Services are being used across the world. There is now a more flexible, and fully featured [Stats API](/reference/api/metrics-stats/historical-stats/) available.
//
// https://developer.fastly.com/reference/api/metrics-stats/stats
type Stats struct {
	sdkConfiguration sdkConfiguration
}

func newStats(sdkConfig sdkConfiguration) *Stats {
	return &Stats{
		sdkConfiguration: sdkConfig,
	}
}

// GetServiceStats - Get stats for a service
// Get the stats from a service for a block of time. This lists all stats by PoP location, starting with AMS. This call requires parameters to select block of time to query. Use either a timestamp range (using start_time and end_time) or a specified month/year combo (using month and year).
func (s *Stats) GetServiceStats(ctx context.Context, request operations.GetServiceStatsRequest) (*operations.GetServiceStatsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/service/{service_id}/stats/summary", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
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

	res := &operations.GetServiceStatsResponse{
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
			var out components.Stats
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Stats = &out
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
