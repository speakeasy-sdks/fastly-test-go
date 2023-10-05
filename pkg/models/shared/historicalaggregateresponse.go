// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// HistoricalAggregateResponseMeta - Meta information about the scope of the query in a human readable format.
type HistoricalAggregateResponseMeta struct {
	By     *string `json:"by,omitempty"`
	From   *string `json:"from,omitempty"`
	Region *string `json:"region,omitempty"`
	To     *string `json:"to,omitempty"`
}

func (o *HistoricalAggregateResponseMeta) GetBy() *string {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *HistoricalAggregateResponseMeta) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *HistoricalAggregateResponseMeta) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *HistoricalAggregateResponseMeta) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}

type HistoricalAggregateResponse struct {
	Data []Results `json:"data,omitempty"`
	// Meta information about the scope of the query in a human readable format.
	Meta *HistoricalAggregateResponseMeta `json:"meta,omitempty"`
	// If the query was not successful, this will provide a string that explains why.
	Msg *string `json:"msg,omitempty"`
	// Whether or not we were able to successfully execute the query.
	Status *string `json:"status,omitempty"`
}

func (o *HistoricalAggregateResponse) GetData() []Results {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *HistoricalAggregateResponse) GetMeta() *HistoricalAggregateResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *HistoricalAggregateResponse) GetMsg() *string {
	if o == nil {
		return nil
	}
	return o.Msg
}

func (o *HistoricalAggregateResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
