// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type TLSConfigurationResponse struct {
	Data *TLSConfigurationResponseData `json:"data,omitempty"`
}

func (o *TLSConfigurationResponse) GetData() *TLSConfigurationResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}