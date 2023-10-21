// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"Fastly/pkg/utils"
	"net/http"
)

type GetHistStatsAggregatedRequest struct {
	// Duration of sample windows. One of:
	//   * `hour` - Group data by hour.
	//   * `minute` - Group data by minute.
	//   * `day` - Group data by day.
	//
	By *shared.By `default:"day" queryParam:"style=form,explode=true,name=by"`
	// Timestamp that defines the start of the window for which to fetch statistics, including the timestamp itself. Accepts Unix timestamps, or any form of input parsable by the [Chronic Ruby library](https://github.com/mojombo/chronic), such as 'yesterday', or 'two weeks ago'. Default varies based on the value of `by`.
	//
	From *string `queryParam:"style=form,explode=true,name=from"`
	// Limit query to a specific geographic region. One of:
	//   * `usa` - North America.
	//   * `europe` - Europe.
	//   * `anzac` - Australia and New Zealand.
	//   * `asia` - Asia.
	//   * `asia_india` - India.
	//   * `asia_southkorea` - South Korea.
	//   * `africa_std` - Africa.
	//   * `southamerica_std` - South America.
	//
	Region *shared.Region `queryParam:"style=form,explode=true,name=region"`
	// Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.
	//
	To *string `default:"now" queryParam:"style=form,explode=true,name=to"`
}

func (g GetHistStatsAggregatedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetHistStatsAggregatedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetHistStatsAggregatedRequest) GetBy() *shared.By {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *GetHistStatsAggregatedRequest) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *GetHistStatsAggregatedRequest) GetRegion() *shared.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *GetHistStatsAggregatedRequest) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}

type GetHistStatsAggregatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	HistoricalAggregateResponse *shared.HistoricalAggregateResponse
}

func (o *GetHistStatsAggregatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHistStatsAggregatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHistStatsAggregatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetHistStatsAggregatedResponse) GetHistoricalAggregateResponse() *shared.HistoricalAggregateResponse {
	if o == nil {
		return nil
	}
	return o.HistoricalAggregateResponse
}
