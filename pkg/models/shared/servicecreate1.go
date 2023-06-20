// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ServiceCreateType - The type of this service.
type ServiceCreateType string

const (
	ServiceCreateTypeVcl  ServiceCreateType = "vcl"
	ServiceCreateTypeWasm ServiceCreateType = "wasm"
)

func (e ServiceCreateType) ToPointer() *ServiceCreateType {
	return &e
}

func (e *ServiceCreateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vcl":
		fallthrough
	case "wasm":
		*e = ServiceCreateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServiceCreateType: %v", v)
	}
}

type ServiceCreate1 struct {
	// A freeform descriptive note.
	Comment *string `form:"name=comment"`
	// Alphanumeric string identifying the customer.
	CustomerID *string `form:"name=customer_id"`
	// The name of the service.
	Name *string `form:"name=name"`
	// The type of this service.
	Type *ServiceCreateType `form:"name=type"`
}
