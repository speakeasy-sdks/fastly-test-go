// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TLSBulkCertificateResponse struct {
	Data *TLSBulkCertificateResponseData `json:"data,omitempty"`
}

func (o *TLSBulkCertificateResponse) GetData() *TLSBulkCertificateResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}
