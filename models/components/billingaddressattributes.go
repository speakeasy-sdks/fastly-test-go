// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type BillingAddressAttributes struct {
	// The first address line.
	Address1 *string `json:"address_1,omitempty"`
	// The second address line.
	Address2 *string `json:"address_2,omitempty"`
	// The city name.
	City *string `json:"city,omitempty"`
	// ISO 3166-1 two-letter country code.
	Country    *string `json:"country,omitempty"`
	CustomerID *string `json:"customer_id,omitempty"`
	// Other locality.
	Locality *string `json:"locality,omitempty"`
	// Postal code (ZIP code for US addresses).
	PostalCode *string `json:"postal_code,omitempty"`
	// The state or province name.
	State *string `json:"state,omitempty"`
}

func (o *BillingAddressAttributes) GetAddress1() *string {
	if o == nil {
		return nil
	}
	return o.Address1
}

func (o *BillingAddressAttributes) GetAddress2() *string {
	if o == nil {
		return nil
	}
	return o.Address2
}

func (o *BillingAddressAttributes) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *BillingAddressAttributes) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *BillingAddressAttributes) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *BillingAddressAttributes) GetLocality() *string {
	if o == nil {
		return nil
	}
	return o.Locality
}

func (o *BillingAddressAttributes) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *BillingAddressAttributes) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}
