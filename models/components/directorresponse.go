// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// DirectorResponseType - What type of load balance group to use.
type DirectorResponseType int64

const (
	DirectorResponseTypeOne   DirectorResponseType = 1
	DirectorResponseTypeThree DirectorResponseType = 3
	DirectorResponseTypeFour  DirectorResponseType = 4
)

func (e DirectorResponseType) ToPointer() *DirectorResponseType {
	return &e
}

func (e *DirectorResponseType) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 3:
		fallthrough
	case 4:
		*e = DirectorResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DirectorResponseType: %v", v)
	}
}

type DirectorResponse struct {
	// List of backends associated to a director.
	Backends []Backend `json:"backends,omitempty"`
	// Unused.
	Capacity *int64 `json:"capacity,omitempty"`
	// A freeform descriptive note.
	Comment *string `json:"comment,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name for the Director.
	Name *string `json:"name,omitempty"`
	// The percentage of capacity that needs to be up for a director to be considered up. `0` to `100`.
	Quorum *int64 `default:"75" json:"quorum"`
	// How many backends to search if it fails.
	Retries   *int64  `default:"5" json:"retries"`
	ServiceID *string `json:"service_id,omitempty"`
	// Selected POP to serve as a shield for the backends. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](/reference/api/utils/pops/) to get a list of available POPs used for shielding.
	Shield *string `default:"null" json:"shield"`
	// What type of load balance group to use.
	Type *DirectorResponseType `default:"1" json:"type"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (d DirectorResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DirectorResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DirectorResponse) GetBackends() []Backend {
	if o == nil {
		return nil
	}
	return o.Backends
}

func (o *DirectorResponse) GetCapacity() *int64 {
	if o == nil {
		return nil
	}
	return o.Capacity
}

func (o *DirectorResponse) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *DirectorResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *DirectorResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *DirectorResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *DirectorResponse) GetQuorum() *int64 {
	if o == nil {
		return nil
	}
	return o.Quorum
}

func (o *DirectorResponse) GetRetries() *int64 {
	if o == nil {
		return nil
	}
	return o.Retries
}

func (o *DirectorResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *DirectorResponse) GetShield() *string {
	if o == nil {
		return nil
	}
	return o.Shield
}

func (o *DirectorResponse) GetType() *DirectorResponseType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *DirectorResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *DirectorResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
