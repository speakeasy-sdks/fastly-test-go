// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MutualAuthenticationInput struct {
	Data *MutualAuthenticationDataInput `json:"data,omitempty"`
}

func (o *MutualAuthenticationInput) GetData() *MutualAuthenticationDataInput {
	if o == nil {
		return nil
	}
	return o.Data
}
