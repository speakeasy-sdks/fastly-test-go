// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetHistStatsFieldSecurity struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

type GetHistStatsFieldRequest struct {
	// Duration of sample windows. One of:
	//   * `hour` - Group data by hour.
	//   * `minute` - Group data by minute.
	//   * `day` - Group data by day.
	//
	By *shared.By `queryParam:"style=form,explode=true,name=by"`
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
	Region *shared.Region `queryParam:"style=form,explode=true,name=region"`
	// Timestamp that defines the end of the window for which to fetch statistics. Accepts the same formats as `from`.
	//
	To *string `queryParam:"style=form,explode=true,name=to"`
}

type GetHistStatsFieldResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	HistoricalFieldResponse *shared.HistoricalFieldResponse
}