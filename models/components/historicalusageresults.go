// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type HistoricalUsageResults struct {
	Bandwidth       *float64 `json:"bandwidth,omitempty"`
	ComputeRequests *float64 `json:"compute_requests,omitempty"`
	Requests        *float64 `json:"requests,omitempty"`
}

func (o *HistoricalUsageResults) GetBandwidth() *float64 {
	if o == nil {
		return nil
	}
	return o.Bandwidth
}

func (o *HistoricalUsageResults) GetComputeRequests() *float64 {
	if o == nil {
		return nil
	}
	return o.ComputeRequests
}

func (o *HistoricalUsageResults) GetRequests() *float64 {
	if o == nil {
		return nil
	}
	return o.Requests
}