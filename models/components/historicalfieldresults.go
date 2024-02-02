// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type HistoricalFieldResults struct {
	AdditionalProperties map[string]string `additionalProperties:"true" json:"-"`
	ServiceID            *string           `json:"service_id,omitempty"`
	StartTime            *string           `json:"start_time,omitempty"`
}

func (h HistoricalFieldResults) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HistoricalFieldResults) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HistoricalFieldResults) GetAdditionalProperties() map[string]string {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *HistoricalFieldResults) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *HistoricalFieldResults) GetStartTime() *string {
	if o == nil {
		return nil
	}
	return o.StartTime
}
