// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type TLSConfiguration struct {
	Data *TLSConfigurationData `json:"data,omitempty"`
}

func (o *TLSConfiguration) GetData() *TLSConfigurationData {
	if o == nil {
		return nil
	}
	return o.Data
}
