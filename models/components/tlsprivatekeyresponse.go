// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type TLSPrivateKeyResponse struct {
	Data *TLSPrivateKeyResponseData `json:"data,omitempty"`
}

func (o *TLSPrivateKeyResponse) GetData() *TLSPrivateKeyResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}
