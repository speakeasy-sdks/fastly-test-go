// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

// TLSSubscriptionResponseAttributesState - The current state of your subscription.
type TLSSubscriptionResponseAttributesState string

const (
	TLSSubscriptionResponseAttributesStatePending    TLSSubscriptionResponseAttributesState = "pending"
	TLSSubscriptionResponseAttributesStateProcessing TLSSubscriptionResponseAttributesState = "processing"
	TLSSubscriptionResponseAttributesStateIssued     TLSSubscriptionResponseAttributesState = "issued"
	TLSSubscriptionResponseAttributesStateRenewing   TLSSubscriptionResponseAttributesState = "renewing"
	TLSSubscriptionResponseAttributesStateFailed     TLSSubscriptionResponseAttributesState = "failed"
)

func (e TLSSubscriptionResponseAttributesState) ToPointer() *TLSSubscriptionResponseAttributesState {
	return &e
}

func (e *TLSSubscriptionResponseAttributesState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "processing":
		fallthrough
	case "issued":
		fallthrough
	case "renewing":
		fallthrough
	case "failed":
		*e = TLSSubscriptionResponseAttributesState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TLSSubscriptionResponseAttributesState: %v", v)
	}
}

type TLSSubscriptionResponseAttributes struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The current state of your subscription.
	State *TLSSubscriptionResponseAttributesState `json:"state,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (t TLSSubscriptionResponseAttributes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSSubscriptionResponseAttributes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSSubscriptionResponseAttributes) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TLSSubscriptionResponseAttributes) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *TLSSubscriptionResponseAttributes) GetState() *TLSSubscriptionResponseAttributesState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *TLSSubscriptionResponseAttributes) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
