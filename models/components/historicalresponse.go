// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// HistoricalResponseMeta - Meta information about the scope of the query in a human readable format.
type HistoricalResponseMeta struct {
	By     *string `json:"by,omitempty"`
	From   *string `json:"from,omitempty"`
	Region *string `json:"region,omitempty"`
	To     *string `json:"to,omitempty"`
}

func (o *HistoricalResponseMeta) GetBy() *string {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *HistoricalResponseMeta) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *HistoricalResponseMeta) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *HistoricalResponseMeta) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}

type HistoricalResponse struct {
	// Contains the results of the query, organized by *service ID*, into arrays where each element describes one service over a *time span*.
	Data map[string][]Results `json:"data,omitempty"`
	// Meta information about the scope of the query in a human readable format.
	Meta *HistoricalResponseMeta `json:"meta,omitempty"`
	// If the query was not successful, this will provide a string that explains why.
	Msg *string `json:"msg,omitempty"`
	// Whether or not we were able to successfully execute the query.
	Status *string `json:"status,omitempty"`
}

func (o *HistoricalResponse) GetData() map[string][]Results {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *HistoricalResponse) GetMeta() *HistoricalResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *HistoricalResponse) GetMsg() *string {
	if o == nil {
		return nil
	}
	return o.Msg
}

func (o *HistoricalResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}