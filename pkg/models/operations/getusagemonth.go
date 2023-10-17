// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"Fastly/pkg/models/shared"
	"net/http"
)

type GetUsageMonthRequest struct {
	// If `true`, return results as billable units.
	BillableUnits *bool `queryParam:"style=form,explode=true,name=billable_units"`
	// 2-digit month.
	Month *string `queryParam:"style=form,explode=true,name=month"`
	// 4-digit year.
	Year *string `queryParam:"style=form,explode=true,name=year"`
}

func (o *GetUsageMonthRequest) GetBillableUnits() *bool {
	if o == nil {
		return nil
	}
	return o.BillableUnits
}

func (o *GetUsageMonthRequest) GetMonth() *string {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *GetUsageMonthRequest) GetYear() *string {
	if o == nil {
		return nil
	}
	return o.Year
}

type GetUsageMonthResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	HistoricalUsageMonthResponse *shared.HistoricalUsageMonthResponse
}

func (o *GetUsageMonthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUsageMonthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUsageMonthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUsageMonthResponse) GetHistoricalUsageMonthResponse() *shared.HistoricalUsageMonthResponse {
	if o == nil {
		return nil
	}
	return o.HistoricalUsageMonthResponse
}
