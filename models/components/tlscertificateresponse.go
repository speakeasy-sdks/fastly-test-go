// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type TLSCertificateResponse struct {
	Data *TLSCertificateResponseData `json:"data,omitempty"`
}

func (o *TLSCertificateResponse) GetData() *TLSCertificateResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}
