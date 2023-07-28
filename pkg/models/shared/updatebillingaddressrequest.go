// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateBillingAddressRequestDataInput struct {
	Attributes *BillingAddressAttributesInput `json:"attributes,omitempty"`
	// Resource type
	Type *TypeBillingAddress `json:"type,omitempty"`
}

func (o *UpdateBillingAddressRequestDataInput) GetAttributes() *BillingAddressAttributesInput {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *UpdateBillingAddressRequestDataInput) GetType() *TypeBillingAddress {
	if o == nil {
		return nil
	}
	return o.Type
}

// UpdateBillingAddressRequestInput - One or more billing address attributes
type UpdateBillingAddressRequestInput struct {
	Data *UpdateBillingAddressRequestDataInput `json:"data,omitempty"`
	// When set to true, the address will be saved without verification
	SkipVerification *bool `json:"skip_verification,omitempty"`
}

func (o *UpdateBillingAddressRequestInput) GetData() *UpdateBillingAddressRequestDataInput {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *UpdateBillingAddressRequestInput) GetSkipVerification() *bool {
	if o == nil {
		return nil
	}
	return o.SkipVerification
}
