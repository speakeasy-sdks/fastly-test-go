// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
	"github.com/speakeasy-sdks/fastly-test-go/models/components"
	"net/http"
)

type GetHistStatsRequest struct {
	// Duration of sample windows. One of:
	//   * `hour` - Group data by hour.
	//   * `minute` - Group data by minute.
	//   * `day` - Group data by day.
	//
	By *components.By `default:"day" queryParam:"style=form,explode=true,name=by"`
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
	Region *components.Region `queryParam:"style=form,explode=true,name=region"`
	// Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.
	//
	To *string `default:"now" queryParam:"style=form,explode=true,name=to"`
}

func (g GetHistStatsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetHistStatsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetHistStatsRequest) GetBy() *components.By {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *GetHistStatsRequest) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *GetHistStatsRequest) GetRegion() *components.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *GetHistStatsRequest) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}

type GetHistStatsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	HistoricalResponse *components.HistoricalResponse
}

func (o *GetHistStatsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHistStatsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHistStatsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetHistStatsResponse) GetHistoricalResponse() *components.HistoricalResponse {
	if o == nil {
		return nil
	}
	return o.HistoricalResponse
}
