// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type MutualAuthenticationResponse struct {
	Data *MutualAuthenticationResponseData `json:"data,omitempty"`
}

func (o *MutualAuthenticationResponse) GetData() *MutualAuthenticationResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}
