// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

type HistoricalServices struct {
	AdditionalProperties map[string]HistoricalUsageResults `additionalProperties:"true" json:"-"`
	// The name of the service.
	Name *string `json:"name,omitempty"`
}

func (h HistoricalServices) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HistoricalServices) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HistoricalServices) GetAdditionalProperties() map[string]HistoricalUsageResults {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *HistoricalServices) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
