// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"Fastly/pkg/utils"
	"net/http"
)

type GetHistStatsRequest struct {
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

func (g GetHistStatsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetHistStatsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetHistStatsRequest) GetBy() *shared.By {
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

func (o *GetHistStatsRequest) GetRegion() *shared.Region {
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
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	HistoricalResponse *shared.HistoricalResponse
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

func (o *GetHistStatsResponse) GetHistoricalResponse() *shared.HistoricalResponse {
	if o == nil {
		return nil
	}
	return o.HistoricalResponse
}
