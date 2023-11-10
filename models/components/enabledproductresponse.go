// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Links struct {
	// Location of resource
	Self *string `json:"self,omitempty"`
	// Location of the service resource
	Service *string `json:"service,omitempty"`
}

func (o *Links) GetSelf() *string {
	if o == nil {
		return nil
	}
	return o.Self
}

func (o *Links) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

type Product struct {
	// Product identifier
	ID *string `json:"id,omitempty"`
	// Name of the object
	Object *string `json:"object,omitempty"`
}

func (o *Product) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Product) GetObject() *string {
	if o == nil {
		return nil
	}
	return o.Object
}

type EnabledProductResponseService struct {
	// Service identifier
	ID *string `json:"id,omitempty"`
	// Name of the object
	Object *string `json:"object,omitempty"`
}

func (o *EnabledProductResponseService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EnabledProductResponseService) GetObject() *string {
	if o == nil {
		return nil
	}
	return o.Object
}

type EnabledProductResponse struct {
	Links   *Links                         `json:"_links,omitempty"`
	Product *Product                       `json:"product,omitempty"`
	Service *EnabledProductResponseService `json:"service,omitempty"`
}

func (o *EnabledProductResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *EnabledProductResponse) GetProduct() *Product {
	if o == nil {
		return nil
	}
	return o.Product
}

func (o *EnabledProductResponse) GetService() *EnabledProductResponseService {
	if o == nil {
		return nil
	}
	return o.Service
}
