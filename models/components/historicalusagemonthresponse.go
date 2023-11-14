// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type HistoricalUsageMonthResponseData struct {
	CustomerID *string                       `json:"customer_id,omitempty"`
	Services   map[string]HistoricalServices `json:"services,omitempty"`
	Total      *HistoricalUsageResults       `json:"total,omitempty"`
}

func (o *HistoricalUsageMonthResponseData) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *HistoricalUsageMonthResponseData) GetServices() map[string]HistoricalServices {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *HistoricalUsageMonthResponseData) GetTotal() *HistoricalUsageResults {
	if o == nil {
		return nil
	}
	return o.Total
}

// HistoricalUsageMonthResponseMeta - Meta information about the scope of the query in a human readable format.
type HistoricalUsageMonthResponseMeta struct {
	By     *string `json:"by,omitempty"`
	From   *string `json:"from,omitempty"`
	Region *string `json:"region,omitempty"`
	To     *string `json:"to,omitempty"`
}

func (o *HistoricalUsageMonthResponseMeta) GetBy() *string {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *HistoricalUsageMonthResponseMeta) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *HistoricalUsageMonthResponseMeta) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *HistoricalUsageMonthResponseMeta) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}

type HistoricalUsageMonthResponse struct {
	Data *HistoricalUsageMonthResponseData `json:"data,omitempty"`
	// Meta information about the scope of the query in a human readable format.
	Meta *HistoricalUsageMonthResponseMeta `json:"meta,omitempty"`
	// If the query was not successful, this will provide a string that explains why.
	Msg *string `json:"msg,omitempty"`
	// Whether or not we were able to successfully execute the query.
	Status *string `json:"status,omitempty"`
}

func (o *HistoricalUsageMonthResponse) GetData() *HistoricalUsageMonthResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *HistoricalUsageMonthResponse) GetMeta() *HistoricalUsageMonthResponseMeta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *HistoricalUsageMonthResponse) GetMsg() *string {
	if o == nil {
		return nil
	}
	return o.Msg
}

func (o *HistoricalUsageMonthResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}