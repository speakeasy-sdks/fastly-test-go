// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// HistoricalFieldResponseMeta - Meta information about the scope of the query in a human readable format.
type HistoricalFieldResponseMeta struct {
	By     *string `json:"by,omitempty"`
	From   *string `json:"from,omitempty"`
	Region *string `json:"region,omitempty"`
	To     *string `json:"to,omitempty"`
}

func (o *HistoricalFieldResponseMeta) GetBy() *string {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *HistoricalFieldResponseMeta) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *HistoricalFieldResponseMeta) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *HistoricalFieldResponseMeta) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}

type HistoricalFieldResponse struct {
	Data map[string][]HistoricalFieldResults `json:"data,omitempty"`
	// Meta information about the scope of the query in a human readable format.
	Meta *HistoricalFieldResponseMeta `json:"meta,omitempty"`
	// If the query was not successful, this will provide a string that explains why.
	Msg *string `json:"msg,omitempty"`
	// Whether or not we were able to successfully execute the query.
	Status *string `json:"status,omitempty"`
}

func (o *HistoricalFieldResponse) GetData() map[string][]HistoricalFieldResults {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *HistoricalFieldResponse) GetMeta() *HistoricalFieldResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *HistoricalFieldResponse) GetMsg() *string {
	if o == nil {
		return nil
	}
	return o.Msg
}

func (o *HistoricalFieldResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
