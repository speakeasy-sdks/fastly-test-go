// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// HistoricalUsageServiceResponseMeta - Meta information about the scope of the query in a human readable format.
type HistoricalUsageServiceResponseMeta struct {
	By     *string `json:"by,omitempty"`
	From   *string `json:"from,omitempty"`
	Region *string `json:"region,omitempty"`
	To     *string `json:"to,omitempty"`
}

func (o *HistoricalUsageServiceResponseMeta) GetBy() *string {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *HistoricalUsageServiceResponseMeta) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *HistoricalUsageServiceResponseMeta) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *HistoricalUsageServiceResponseMeta) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}

// HistoricalUsageServiceResponse - OK
type HistoricalUsageServiceResponse struct {
	Data *HistoricalUsageResults `json:"data,omitempty"`
	// Meta information about the scope of the query in a human readable format.
	Meta *HistoricalUsageServiceResponseMeta `json:"meta,omitempty"`
	// If the query was not successful, this will provide a string that explains why.
	Msg *string `json:"msg,omitempty"`
	// Whether or not we were able to successfully execute the query.
	Status *string `json:"status,omitempty"`
}

func (o *HistoricalUsageServiceResponse) GetData() *HistoricalUsageResults {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *HistoricalUsageServiceResponse) GetMeta() *HistoricalUsageServiceResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *HistoricalUsageServiceResponse) GetMsg() *string {
	if o == nil {
		return nil
	}
	return o.Msg
}

func (o *HistoricalUsageServiceResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
