// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BillingAddressAttributesInput struct {
	// The first address line.
	Address1 *string `json:"address_1,omitempty"`
	// The second address line.
	Address2 *string `json:"address_2,omitempty"`
	// The city name.
	City *string `json:"city,omitempty"`
	// ISO 3166-1 two-letter country code.
	Country *string `json:"country,omitempty"`
	// Other locality.
	Locality *string `json:"locality,omitempty"`
	// Postal code (ZIP code for US addresses).
	PostalCode *string `json:"postal_code,omitempty"`
	// The state or province name.
	State *string `json:"state,omitempty"`
}

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