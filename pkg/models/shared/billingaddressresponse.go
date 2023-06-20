// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BillingAddressResponseData struct {
	Attributes *BillingAddressAttributes `json:"attributes,omitempty"`
	// Alphanumeric string identifying the billing address.
	ID            *string               `json:"id,omitempty"`
	Relationships *RelationshipCustomer `json:"relationships,omitempty"`
	// Resource type
	Type *TypeBillingAddress `json:"type,omitempty"`
}

// BillingAddressResponse - OK
type BillingAddressResponse struct {
	Data *BillingAddressResponseData `json:"data,omitempty"`
}
