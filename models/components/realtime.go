// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Realtime struct {
	// How long the system will wait before aggregating messages for each second. The most recent data returned will have happened at the moment of the request, minus the aggregation delay.
	AggregateDelay *int64 `json:"AggregateDelay,omitempty"`
	// A list of [records](#record-data-model), each representing one second of time.
	Data []RealtimeEntry `json:"Data,omitempty"`
	// Value to use for subsequent requests.
	Timestamp *int64 `json:"Timestamp,omitempty"`
}

func (o *Realtime) GetAggregateDelay() *int64 {
	if o == nil {
		return nil
	}
	return o.AggregateDelay
}

func (o *Realtime) GetData() []RealtimeEntry {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *Realtime) GetTimestamp() *int64 {
	if o == nil {
		return nil
	}
	return o.Timestamp
}
