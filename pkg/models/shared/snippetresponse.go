// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// SnippetResponseDynamic - Sets the snippet version.
type SnippetResponseDynamic int64

const (
	SnippetResponseDynamicZero SnippetResponseDynamic = 0
	SnippetResponseDynamicOne  SnippetResponseDynamic = 1
)

func (e SnippetResponseDynamic) ToPointer() *SnippetResponseDynamic {
	return &e
}

func (e *SnippetResponseDynamic) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = SnippetResponseDynamic(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SnippetResponseDynamic: %v", v)
	}
}

// SnippetResponseType - The location in generated VCL where the snippet should be placed.
type SnippetResponseType string

const (
	SnippetResponseTypeInit    SnippetResponseType = "init"
	SnippetResponseTypeRecv    SnippetResponseType = "recv"
	SnippetResponseTypeHash    SnippetResponseType = "hash"
	SnippetResponseTypeHit     SnippetResponseType = "hit"
	SnippetResponseTypeMiss    SnippetResponseType = "miss"
	SnippetResponseTypePass    SnippetResponseType = "pass"
	SnippetResponseTypeFetch   SnippetResponseType = "fetch"
	SnippetResponseTypeError   SnippetResponseType = "error"
	SnippetResponseTypeDeliver SnippetResponseType = "deliver"
	SnippetResponseTypeLog     SnippetResponseType = "log"
	SnippetResponseTypeNone    SnippetResponseType = "none"
)

func (e SnippetResponseType) ToPointer() *SnippetResponseType {
	return &e
}

func (e *SnippetResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "init":
		fallthrough
	case "recv":
		fallthrough
	case "hash":
		fallthrough
	case "hit":
		fallthrough
	case "miss":
		fallthrough
	case "pass":
		fallthrough
	case "fetch":
		fallthrough
	case "error":
		fallthrough
	case "deliver":
		fallthrough
	case "log":
		fallthrough
	case "none":
		*e = SnippetResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SnippetResponseType: %v", v)
	}
}

type SnippetResponse struct {
	// The VCL code that specifies exactly what the snippet does.
	Content *string `json:"content,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Sets the snippet version.
	Dynamic *SnippetResponseDynamic `json:"dynamic,omitempty"`
	ID      *string                 `json:"id,omitempty"`
	// The name for the snippet.
	Name *string `json:"name,omitempty"`
	// Priority determines execution order. Lower numbers execute first.
	Priority  *string `json:"priority,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	// The location in generated VCL where the snippet should be placed.
	Type *SnippetResponseType `json:"type,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// String representing the number identifying a version of the service.
	Version *string `json:"version,omitempty"`
}

func (o *SnippetResponse) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *SnippetResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SnippetResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *SnippetResponse) GetDynamic() *SnippetResponseDynamic {
	if o == nil {
		return nil
	}
	return o.Dynamic
}

func (o *SnippetResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SnippetResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SnippetResponse) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *SnippetResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *SnippetResponse) GetType() *SnippetResponseType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SnippetResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *SnippetResponse) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
