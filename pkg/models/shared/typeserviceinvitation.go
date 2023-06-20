// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TypeServiceInvitation - Resource type
type TypeServiceInvitation string

const (
	TypeServiceInvitationServiceInvitation TypeServiceInvitation = "service_invitation"
)

func (e TypeServiceInvitation) ToPointer() *TypeServiceInvitation {
	return &e
}

func (e *TypeServiceInvitation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_invitation":
		*e = TypeServiceInvitation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TypeServiceInvitation: %v", v)
	}
}
