// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type VersionCreateResponse struct {
	Number    *int64  `json:"number,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
}

func (o *VersionCreateResponse) GetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *VersionCreateResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}
