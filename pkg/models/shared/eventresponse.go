// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EventResponse - OK
type EventResponse struct {
	Data *Event `json:"data,omitempty"`
}

func (o *EventResponse) GetData() *Event {
	if o == nil {
		return nil
	}
	return o.Data
}
