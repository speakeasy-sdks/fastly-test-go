// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BillingAddressVerificationErrorResponseErrors struct {
	Candidates []BillingAddressAttributes `json:"candidates,omitempty"`
	Detail     string                     `json:"detail"`
	Status     float64                    `json:"status"`
	Title      string                     `json:"title"`
	// The error type.
	Type string `json:"type"`
}

func (o *BillingAddressVerificationErrorResponseErrors) GetCandidates() []BillingAddressAttributes {
	if o == nil {
		return nil
	}
	return o.Candidates
}

func (o *BillingAddressVerificationErrorResponseErrors) GetDetail() string {
	if o == nil {
		return ""
	}
	return o.Detail
}

func (o *BillingAddressVerificationErrorResponseErrors) GetStatus() float64 {
	if o == nil {
		return 0.0
	}
	return o.Status
}

func (o *BillingAddressVerificationErrorResponseErrors) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *BillingAddressVerificationErrorResponseErrors) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

type BillingAddressVerificationErrorResponse struct {
	Errors []BillingAddressVerificationErrorResponseErrors `json:"errors,omitempty"`
}

func (o *BillingAddressVerificationErrorResponse) GetErrors() []BillingAddressVerificationErrorResponseErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}
