// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/internal/utils"
	"Fastly/models/components"
	"net/http"
)

type GetHistStatsFieldRequest struct {
	// Duration of sample windows. One of:
	//   * `hour` - Group data by hour.
	//   * `minute` - Group data by minute.
	//   * `day` - Group data by day.
	//
	By *components.By `default:"day" queryParam:"style=form,explode=true,name=by"`
	// Name of the stats field.
	Field string `pathParam:"style=simple,explode=false,name=field"`
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

func (g GetHistStatsFieldRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetHistStatsFieldRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetHistStatsFieldRequest) GetBy() *components.By {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *GetHistStatsFieldRequest) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *GetHistStatsFieldRequest) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *GetHistStatsFieldRequest) GetRegion() *components.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *GetHistStatsFieldRequest) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}

type GetHistStatsFieldResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	HistoricalFieldResponse *components.HistoricalFieldResponse
}

func (o *GetHistStatsFieldResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHistStatsFieldResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHistStatsFieldResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetHistStatsFieldResponse) GetHistoricalFieldResponse() *components.HistoricalFieldResponse {
	if o == nil {
		return nil
	}
	return o.HistoricalFieldResponse
}
