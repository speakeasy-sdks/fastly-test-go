// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// StoreResponse - Created
type StoreResponse struct {
	// ID of the store.
	ID *string `json:"id,omitempty"`
	// A human-readable name for the store.
	Name *string `json:"name,omitempty"`
}

func (o *StoreResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *StoreResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
